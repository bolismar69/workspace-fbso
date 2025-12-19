Formato do input `services` para `java-ci.yml`

O workflow principal `java-ci.yml` detecta serviços Java alterados e passa uma string JSON para o workflow reutilizável `callables/build-java-service.yml`.

Formato esperado (JSON array):
[
  {
    "path": "backend/java/quarkus/ms-user-auth",
    "goals": "clean install"
  },
  {
    "path": "backend/java/spring/ms-payment",
    "goals": "clean install -DskipTests"
  }
]

Como disparar manualmente pelo UI do GitHub:
- Vá em Actions → escolha o workflow `Java CI/CD Trigger` (java-ci.yml) → Run workflow → no campo `services` cole o JSON de exemplo acima e execute.

Como disparar via API (exemplo curl):

curl -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer <TOKEN>" \
  https://api.github.com/repos/<OWNER>/<REPO>/actions/workflows/java-ci.yml/dispatches \
  -d '{"ref":"main","inputs":{"services":"[ { \"path\": \"backend/java/quarkus/ms-user-auth\", \"goals\": \"clean install\" } ]"}}'

Observações:
- O campo `services` é uma string contendo JSON (escape necessário se usar curl). No GitHub UI isso não é necessário.
- Certifique-se de que as variáveis em `.github/docker.env` e `.github/docker.java.env` existem e contém `DOCKER_REGISTRY`, `DOCKER_ORG`, `DOCKER_TAG_PREFIX` e outras usadas pelo build.

Publicar mudanças rapidamente
- Script: `scripts/publish.sh`
- Uso: torne-o executável e rode com a mensagem de commit como parâmetro:

```bash
chmod +x scripts/publish.sh
./scripts/publish.sh "UPD: repositorio atualizado ..."
```

O script faz `git add -A`, `git commit -m "..."` e `git push origin main`. Se a mensagem não for fornecida o script exibirá erro e não fará commit.
