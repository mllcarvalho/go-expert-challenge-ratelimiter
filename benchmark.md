# Go Expert Challenge - Avaliação de Rate Limiter

Realizamos testes de carga utilizando o Grafana k6 para analisar o desempenho da solução desenvolvida sob condições de alta demanda. Em todos os experimentos, a aplicação estava rodando dentro de um container Docker, acessível na porta 8080, enquanto o k6 foi executado a partir de um container separado.

## Hardware

- **CPU**: Apple M2
- **RAM**: 8GB DDR4
- **OS**: MacOS Sonoma
- **Runtime**: Docker (alpine)

## Smoke test

Teste com fim de validar se o serviço está respondendo corretamente.

- **Target**: 5 usuários
- **Duração**: 1 minuto

Comando para execução:

```sh
make test_k6_smoke
```

### Resultado

```plaintext
K6_WEB_DASHBOARD=true K6_WEB_DASHBOARD_EXPORT=scripts/k6/smoke-test-report.html k6 run scripts/k6/smoke.test.js


         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: /scripts/smoke/smoke.test.js
 web dashboard: http://127.0.0.1:5665
        output: -

     scenarios: (100.00%) 1 scenario, 5 max VUs, 1m30s max duration (incl. graceful stop):
              * default: 5 looping VUs for 1m0s (gracefulStop: 30s)


     data_received..................: 52 MB  864 kB/s
     data_sent......................: 15 MB  250 kB/s
     http_req_blocked...............: min=458ns      med=875ns    avg=1.78µs  p(90)=2.5µs    p(95)=3.25µs  max=3.34ms   count=152786
     http_req_connecting............: min=0s         med=0s       avg=24ns    p(90)=0s       p(95)=0s      max=807.95µs count=152786
     http_req_duration..............: min=-2293973ns med=347.41µs avg=1.93ms  p(90)=753.08µs p(95)=1.2ms   max=2.56s    count=152786
       { expected_response:true }...: min=199.16µs   med=525.08µs avg=1.12ms  p(90)=1.16ms   p(95)=1.75ms  max=54.72ms  count=660   
     ✓ { status:200 }...............: min=199.16µs   med=525.08µs avg=1.12ms  p(90)=1.16ms   p(95)=1.75ms  max=54.72ms  count=660   
     ✓ { status:429 }...............: min=-2293973ns med=346.91µs avg=1.93ms  p(90)=749.25µs p(95)=1.2ms   max=2.56s    count=152126
     ✓ { status:500 }...............: min=0s         med=0s       avg=0s      p(90)=0s       p(95)=0s      max=0s       count=0     
     http_req_failed................: 99.56% 152126 out of 152786
     http_req_receiving.............: min=3.75µs     med=8.12µs   avg=16.36µs p(90)=30.66µs  p(95)=41.75µs max=9ms      count=152786
     http_req_sending...............: min=-2505681ns med=2.04µs   avg=4.18µs  p(90)=6.87µs   p(95)=9µs     max=9.06ms   count=152786
     http_req_tls_handshaking.......: min=0s         med=0s       avg=0s      p(90)=0s       p(95)=0s      max=0s       count=152786
     http_req_waiting...............: min=0s         med=329.95µs avg=1.91ms  p(90)=723.66µs p(95)=1.17ms  max=2.56s    count=152786
     http_reqs......................: 152786 2546.472896/s
     iteration_duration.............: min=243.75µs   med=763.95µs avg=3.92ms  p(90)=1.68ms   p(95)=2.18ms  max=2.56s    count=76393 
     iterations.....................: 76393  1273.236448/s
     vus............................: 5      min=5                max=5
     vus_max........................: 5      min=5                max=5


running (1m00.0s), 0/5 VUs, 76393 complete and 0 interrupted iterations
default ✓ [======================================] 5 VUs  1m0s
```

## Stress test

Teste com fim de validar se o serviço está respondendo corretamente sob pressão.

- Stage 1:
    - **Target**: 30 usuários
    - **Duração**: 10 minutos
- Stage 2:
    - **Target**: 60 usuários
    - **Duração**: 10 minutos
- Stage 3:
    - **Target**: 0 usuários
    - **Duração**: 5 minutos

Comando para execução:

```sh
make test_k6_stress
```

### Resultado

```plaintext
K6_WEB_DASHBOARD=true K6_WEB_DASHBOARD_EXPORT=scripts/k6/stress/stress-test-report.html k6 run scripts/k6/stress/stress.test.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: /scripts/stress/stress.test.js
 web dashboard: http://127.0.0.1:5665
        output: -

     scenarios: (100.00%) 1 scenario, 60 max VUs, 25m30s max duration (incl. graceful stop):
              * default: Up to 60 looping VUs for 25m0s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


     data_received..................: 29 MB  19 kB/s
     data_sent......................: 8.7 MB 5.8 kB/s
     http_req_blocked...............: min=583ns    med=6.2µs   avg=11.4µs  p(90)=18.62µs  p(95)=23.15µs  max=10.24ms count=88924
     http_req_connecting............: min=0s       med=0s      avg=907ns   p(90)=0s       p(95)=0s       max=9.22ms  count=88924
     http_req_duration..............: min=214.2µs  med=2.29ms  avg=2.68ms  p(90)=4.48ms   p(95)=5.86ms   max=31.09ms count=88924
       { expected_response:true }...: min=377.25µs med=1.72ms  avg=2.19ms  p(90)=3.7ms    p(95)=4.72ms   max=31.09ms count=15136
     ✓ { status:200 }...............: min=377.25µs med=1.72ms  avg=2.19ms  p(90)=3.7ms    p(95)=4.72ms   max=31.09ms count=15136
     ✓ { status:429 }...............: min=214.2µs  med=2.41ms  avg=2.78ms  p(90)=4.61ms   p(95)=6.02ms   max=30.17ms count=73788
     ✓ { status:500 }...............: min=0s       med=0s      avg=0s      p(90)=0s       p(95)=0s       max=0s      count=0    
     http_req_failed................: 82.97% 73788 out of 88924
     http_req_receiving.............: min=5.12µs   med=39.95µs avg=59.99µs p(90)=111.08µs p(95)=140.54µs max=8.53ms  count=88924
     http_req_sending...............: min=1.54µs   med=15.75µs avg=29.3µs  p(90)=56.37µs  p(95)=68.45µs  max=18.29ms count=88924
     http_req_tls_handshaking.......: min=0s       med=0s      avg=0s      p(90)=0s       p(95)=0s       max=0s      count=88924
     http_req_waiting...............: min=192.08µs med=2.18ms  avg=2.59ms  p(90)=4.38ms   p(95)=5.72ms   max=31.02ms count=88924
     http_reqs......................: 88924  59.268509/s
     iteration_duration.............: min=1s       med=1s      avg=1s      p(90)=1.01s    p(95)=1.01s    max=1.03s   count=44462
     iterations.....................: 44462  29.634255/s
     vus............................: 1      min=1              max=60
     vus_max........................: 60     min=60             max=60


running (25m00.4s), 00/60 VUs, 44462 complete and 0 interrupted iterations
default ✓ [======================================] 00/60 VUs  25m0s
```