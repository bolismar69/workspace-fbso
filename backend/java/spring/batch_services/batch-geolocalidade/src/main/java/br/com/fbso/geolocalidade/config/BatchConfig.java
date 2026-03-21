package br.com.fbso.geolocalidade.config;

import br.com.fbso.geolocalidade.dto.DistritoCsvDTO;
import br.com.fbso.geolocalidade.dto.MunicipioCsvDTO;
import br.com.fbso.geolocalidade.dto.SubdistritoCsvDTO;
import br.com.fbso.geolocalidade.entity.Distrito;
import br.com.fbso.geolocalidade.entity.Municipio;
import br.com.fbso.geolocalidade.entity.Subdistrito;
import br.com.fbso.geolocalidade.processor.DistritoProcessor;
import br.com.fbso.geolocalidade.processor.MunicipioProcessor;
import br.com.fbso.geolocalidade.processor.SubdistritoProcessor;
import br.com.fbso.geolocalidade.repository.DistritoRepository;
import br.com.fbso.geolocalidade.repository.MunicipioRepository;
import br.com.fbso.geolocalidade.repository.SubdistritoRepository;
import org.springframework.batch.core.Job;
import org.springframework.batch.core.job.builder.JobBuilder;
import org.springframework.batch.core.repository.JobRepository;
import org.springframework.batch.core.Step;
import org.springframework.batch.core.step.builder.StepBuilder;
import org.springframework.batch.item.ItemProcessor;
import org.springframework.batch.item.ItemReader;
import org.springframework.batch.item.ItemWriter;
import org.springframework.batch.item.data.RepositoryItemWriter;
import org.springframework.batch.item.data.builder.RepositoryItemWriterBuilder;
import org.springframework.batch.item.file.FlatFileItemReader;
import org.springframework.batch.item.file.builder.FlatFileItemReaderBuilder;
import org.springframework.batch.item.file.mapping.DefaultLineMapper;
import org.springframework.batch.item.file.transform.DelimitedLineTokenizer;
import org.springframework.batch.item.file.transform.FieldSet;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.FileSystemResource;
import org.springframework.transaction.PlatformTransactionManager;

@Configuration
public class BatchConfig {

  @Value("${app.import.path}")
  private String importPath;

  @Value("${app.import.files.municipios:DTB_Municipios.csv}")
  private String municipiosFile;

  @Value("${app.import.files.distritos:DTB_Distritos.csv}")
  private String distritosFile;

  @Value("${app.import.files.subdistritos:DTB_Subdistritos.csv}")
  private String subdistritosFile;
  
  private final JobRepository jobRepository;
  private final PlatformTransactionManager transactionManager;

  public BatchConfig(JobRepository jobRepository, PlatformTransactionManager transactionManager) {
    this.jobRepository = jobRepository;
    this.transactionManager = transactionManager;
  }

  // --- JOB DEFINITION ---

  @Bean
  public Job importacaoGeolocalidadeJob(Step stepMunicipio, Step stepDistrito, Step stepSubdistrito) {
    return new JobBuilder("importacaoGeolocalidadeJob", jobRepository)
      .start(stepMunicipio)      // Primeiro as UFs, Regiões e Municípios
      .next(stepDistrito)       // Depois os Distritos (Depende de Município)
      .next(stepSubdistrito)    // Por fim Subdistritos (Depende de Distrito)
      .build();
  }

  // --- STEP 1: MUNICÍPIOS (E HIERARQUIA SUPERIOR) ---

  // --- STEP 1: MUNICÍPIOS ---
  @Bean
  public Step stepMunicipio(
          ItemReader<MunicipioCsvDTO> municipioReader,
          ItemProcessor<MunicipioCsvDTO, Municipio> municipioProcessor,
          ItemWriter<Municipio> municipioWriter) {
    return new StepBuilder("stepMunicipio", jobRepository)
      .<MunicipioCsvDTO, Municipio>chunk(100, transactionManager)
      .reader(municipioReader)
      .processor(municipioProcessor)
      .writer(municipioWriter)
      .build();
  }

  // --- STEP 2: DISTRITOS ---
  @Bean
  public Step stepDistrito(FlatFileItemReader<DistritoCsvDTO> reader,
                            DistritoProcessor processor,
                            RepositoryItemWriter<Distrito> writer) {
    return new StepBuilder("stepDistrito", jobRepository)
      .<DistritoCsvDTO, Distrito>chunk(100, transactionManager)
      .reader(reader)
      .processor(processor)
      .writer(writer)
      .build();
  }

  // --- STEP 3: SUBDISTRITOS ---
  @Bean
  public Step stepSubdistrito(FlatFileItemReader<SubdistritoCsvDTO> reader,
                              SubdistritoProcessor processor,
                              RepositoryItemWriter<Subdistrito> writer) {
    return new StepBuilder("stepSubdistrito", jobRepository)
      .<SubdistritoCsvDTO, Subdistrito>chunk(100, transactionManager)
      .reader(reader)
      .processor(processor)
      .writer(writer)
      .build();
  }

  // --- READER: CONFIGURADO PARA O VOLUME K8S ---

  @Bean
  public FlatFileItemReader<MunicipioCsvDTO> municipioReader() {
    // O caminho vem da variável de ambiente APP_IMPORT_PATH definida no README/K8S
    // String path = System.getenv("APP_IMPORT_PATH") + "/RELATORIO_DTB_BRASIL_2024_MUNICIPIOS.csv";
    
    DelimitedLineTokenizer tokenizer = new DelimitedLineTokenizer();
    tokenizer.setDelimiter(",");
    tokenizer.setStrict(false); // tolera colunas extras (ex.: vírgula final)
    tokenizer.setNames(
      "ufId", "ufNome",
      "regiaoInterId", "regiaoInterNome",
      "regiaoImedId", "regiaoImedNome",
      "municipioCodCurto", "municipioIdCompleto", "municipioNome"
    );

    DefaultLineMapper<MunicipioCsvDTO> lineMapper = new DefaultLineMapper<>();
    lineMapper.setLineTokenizer(tokenizer);
    lineMapper.setFieldSetMapper((FieldSet fs) -> new MunicipioCsvDTO(
      fs.readString(0),
      fs.readString(1),
      fs.readString(2),
      fs.readString(3),
      fs.readString(4),
      fs.readString(5),
      fs.readString(6),
      fs.readString(7),
      fs.readString(8)
    ));
    lineMapper.afterPropertiesSet();

    return new FlatFileItemReaderBuilder<MunicipioCsvDTO>()
      .name("municipioReader")
      .resource(new FileSystemResource(importPath + "/" + municipiosFile))
      .linesToSkip(7) // Pula os metadados do IBGE
      .lineMapper(lineMapper)
      .encoding("UTF-8")
      .build();
  }

  @Bean
  public FlatFileItemReader<DistritoCsvDTO> distritoReader() {
    DelimitedLineTokenizer tokenizer = new DelimitedLineTokenizer();
    tokenizer.setDelimiter(",");
    tokenizer.setStrict(false);
    tokenizer.setNames(
      "ufId", "ufNome",
      "regInterId", "regInterNome",
      "regImedId", "regImedNome",
      "munCod", "municipioId", "munNome",
      "distritoCodCurto", "distritoIdCompleto", "distritoNome"
    );

    DefaultLineMapper<DistritoCsvDTO> lineMapper = new DefaultLineMapper<>();
    lineMapper.setLineTokenizer(tokenizer);
    lineMapper.setFieldSetMapper((FieldSet fs) -> new DistritoCsvDTO(
      fs.readString(0),
      fs.readString(1),
      fs.readString(2),
      fs.readString(3),
      fs.readString(4),
      fs.readString(5),
      fs.readString(6),
      fs.readString(7),
      fs.readString(8),
      fs.readString(9),
      fs.readString(10),
      fs.readString(11)
    ));
    lineMapper.afterPropertiesSet();

    return new FlatFileItemReaderBuilder<DistritoCsvDTO>()
      .name("distritoReader")
      .resource(new FileSystemResource(importPath + "/" + distritosFile))
      .linesToSkip(7)
      .lineMapper(lineMapper)
      .encoding("UTF-8")
      .build();
  }

  @Bean
  public FlatFileItemReader<SubdistritoCsvDTO> subdistritoReader() {
    DelimitedLineTokenizer tokenizer = new DelimitedLineTokenizer();
    tokenizer.setDelimiter(",");
    tokenizer.setStrict(false);
    tokenizer.setNames(
      "ufId", "ufNome",
      "regInterId", "regInterNome",
      "regImedId", "regImedNome",
      "munCod", "municipioId", "munNome",
      "distCod", "distritoId", "distNome",
      "subdistritoCodCurto", "subdistritoIdCompleto", "subdistritoNome"
    );

    DefaultLineMapper<SubdistritoCsvDTO> lineMapper = new DefaultLineMapper<>();
    lineMapper.setLineTokenizer(tokenizer);
    lineMapper.setFieldSetMapper((FieldSet fs) -> new SubdistritoCsvDTO(
      fs.readString(0),
      fs.readString(1),
      fs.readString(2),
      fs.readString(3),
      fs.readString(4),
      fs.readString(5),
      fs.readString(6),
      fs.readString(7),
      fs.readString(8),
      fs.readString(9),
      fs.readString(10),
      fs.readString(11),
      fs.readString(12),
      fs.readString(13),
      fs.readString(14)
    ));
    lineMapper.afterPropertiesSet();

    return new FlatFileItemReaderBuilder<SubdistritoCsvDTO>()
      .name("subdistritoReader")
      .resource(new FileSystemResource(importPath + "/" + subdistritosFile))
      .linesToSkip(7)
      .lineMapper(lineMapper)
      .encoding("UTF-8")
      .build();
  }

  // WRITER CONFIGURADO PARA USAR O REPOSITÓRIO SPRING DATA JPA

  @Bean
  public RepositoryItemWriter<Municipio> municipioWriter(MunicipioRepository repository) {  
    return new RepositoryItemWriterBuilder<Municipio>()
      .repository(repository)
      .methodName("save")
      .build();
  }

  @Bean
  public RepositoryItemWriter<Distrito> distritoWriter(DistritoRepository repository) {
    return new RepositoryItemWriterBuilder<Distrito>()
      .repository(repository)
      .methodName("save")
      .build();
    }

  @Bean
  public RepositoryItemWriter<Subdistrito> subdistritoWriter(SubdistritoRepository repository) {
    return new RepositoryItemWriterBuilder<Subdistrito>()
      .repository(repository)
      .methodName("save")
      .build();
  }
}