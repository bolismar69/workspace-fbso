package com.fbso.queue;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Component;

@Component
public class SampleQueueListener {

    // Configuração do nome da fila deve vir do application.yml
    @RabbitListener(queues = "${fbso.queue.name}")
    public void processMessage(String message) {
        // Implementar Clean Architecture: chamar o Caso de Uso
        System.out.println("Mensagem recebida da fila: " + message);
    }
}
