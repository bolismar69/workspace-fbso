package com.fbso.transaction.queue.listener;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Este componente lida com mensagens de uma fila de transações
@Component
public class TransactionProcessorListener {

    private static final Logger log = LoggerFactory.getLogger(TransactionProcessorListener.class);

    // Usa a configuração do application.yml para o nome da fila
    @RabbitListener(queues = "${fbso.transaction.queue.name}")
    public void processTransaction(String transactionJson) {
        
        log.info("Processando nova transação. Tamanho da mensagem: {} bytes", transactionJson.length());
        
        // **Fase de Clean Architecture:** // Aqui, a mensagem JSON deve ser validada e passada para um Caso de Uso (Use Case)
        // Exemplo: transactionUseCase.execute(transactionJson);

        try {
             // Simulação de processamento
             Thread.sleep(100); 
             log.info("Transação processada com sucesso.");
        } catch (InterruptedException e) {
             Thread.currentThread().interrupt();
             log.error("Erro ao processar transação: {}", e.getMessage());
        }
    }
}
