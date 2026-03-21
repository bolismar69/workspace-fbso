package br.com.fbso.geolocalidade.load;

import br.com.fbso.geolocalidade.repository.DistritoRepository;
import br.com.fbso.geolocalidade.repository.MunicipioRepository;
import br.com.fbso.geolocalidade.repository.SubdistritoRepository;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.time.Duration;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.batch.core.BatchStatus;
import org.springframework.batch.core.Job;
import org.springframework.batch.core.JobExecution;
import org.springframework.batch.core.JobParameters;
import org.springframework.batch.core.JobParametersBuilder;
import org.springframework.batch.core.launch.JobLauncher;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.stereotype.Component;

@Component
@ConditionalOnProperty(name = "app.loadtest.enabled", havingValue = "true")
public class LoadTestRunner implements ApplicationRunner {

  private static final Logger log = LoggerFactory.getLogger(LoadTestRunner.class);

  private final JobLauncher jobLauncher;
  private final Job job;
  private final MunicipioRepository municipioRepository;
  private final DistritoRepository distritoRepository;
  private final SubdistritoRepository subdistritoRepository;
  private final ConfigurableApplicationContext applicationContext;

  @Value("${app.import.path}")
  private String importPath;

  @Value("${app.import.files.municipios:DTB_Municipios.csv}")
  private String municipiosFile;

  @Value("${app.import.files.distritos:DTB_Distritos.csv}")
  private String distritosFile;

  @Value("${app.import.files.subdistritos:DTB_Subdistritos.csv}")
  private String subdistritosFile;

  public LoadTestRunner(
      JobLauncher jobLauncher,
      @Qualifier("importacaoGeolocalidadeJob") Job job,
      MunicipioRepository municipioRepository,
      DistritoRepository distritoRepository,
      SubdistritoRepository subdistritoRepository,
      ConfigurableApplicationContext applicationContext) {
    this.jobLauncher = jobLauncher;
    this.job = job;
    this.municipioRepository = municipioRepository;
    this.distritoRepository = distritoRepository;
    this.subdistritoRepository = subdistritoRepository;
    this.applicationContext = applicationContext;
  }

  @Override
  public void run(ApplicationArguments args) throws Exception {
    log.info("Load test enabled. importPath={}", importPath);

    logImportFile("municipios", municipiosFile);
    logImportFile("distritos", distritosFile);
    logImportFile("subdistritos", subdistritosFile);

    long startNanos = System.nanoTime();

    JobParameters jobParameters = new JobParametersBuilder()
        .addLong("run.id", System.currentTimeMillis())
        .toJobParameters();

    JobExecution execution = jobLauncher.run(job, jobParameters);

    Duration duration = Duration.ofNanos(System.nanoTime() - startNanos);

    long municipios = municipioRepository.count();
    long distritos = distritoRepository.count();
    long subdistritos = subdistritoRepository.count();

    log.info("Load test finished. status={}, exitStatus={}, durationMs={}, municipios={}, distritos={}, subdistritos={}",
        execution.getStatus(), execution.getExitStatus().getExitCode(), duration.toMillis(), municipios, distritos, subdistritos);

    int code = execution.getStatus() == BatchStatus.COMPLETED ? 0 : 1;
    int exitCode = SpringApplication.exit(applicationContext, () -> code);
    System.exit(exitCode);
  }

  private void logImportFile(String label, String fileName) {
    Path path = Paths.get(importPath, fileName).toAbsolutePath().normalize();
    log.info("IBGE file [{}] path={}, exists={}", label, path, Files.exists(path));
  }
}
