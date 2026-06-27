# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Proje özeti
CDP/martech temalı, tek bir birleşik sistem (Insider mülakat hazırlığı + AI engineering
capstone'u). Dört gevşek bağlı modül, ortak domain: müşteri / segment / etkinlik.

## Roadmap
Güncel 12 haftalık plan ve ilerleme durumu **PLAN.md**'de tutulur. Her oturumun
başında önce PLAN.md'ye bak, hangi haftadayız oradan teyit et. Bu dosyaya
(CLAUDE.md) execution plan veya checklist eklenmez — sadece kalıcı gerçekler.

## Kullanıcının gerçek seviyesi (kalibre edilmiş — varsayım değil)
- **Laravel/PHP:** Sıfır. İlk kez öğreniyor. Adım adım anlat, kısayol/jargon varsayma.
- **Docker/Kubernetes:** Sertifikası var ama kalibrasyon testinde temel sorulara
  (multi-stage build, restart-loop debug, Pod/Service/Deployment farkı, rollback)
  cevap veremedi. Tamamen yeniymiş gibi davran, sertifikaya güvenme.
- **AWS:** Kısmi. S3, Lambda/event trigger, Elastic IP sağlam. VPC/subnet/NAT ve
  EC2 IAM instance-profile pattern'inde net boşluk var.
- **MCP / Claude API:** Terim ve kelime hazinesi biliniyor (sertifikalardan:
  stop_reason, tool_use) ama mekanizma bilinmiyor — tools/resources/prompts
  primitiflerini bilmiyor, tool_use akışını (kim tool'u çalıştırır, tool_result
  nasıl geri gönderilir) karıştırıyor. Prompt caching'i semantic/response cache
  ile karıştırıyor — ikisi farklı.
- **Go:** Sıfır. Ama Insider ilanlarında PHP'ye **eşdeğer ağırlıkta** bir backend
  dili (defalarca "PHP ve Go veya Node.js" deniyor) — "giriş seviyesi ekstra"
  gibi davranma, Laravel API ile aynı kalite barını uygula.
- **Node.js/TypeScript/React/React Native/Flutter/Dart/Java/C#:** Çok yıllı
  production deneyimi var (Netelsan + freelance). Yeni kavramları buna analoji
  kurarak anlat — örn. Python `asyncio` ↔ Node.js event loop / Dart `async/await`.
- Yeni bir konuya geçmeden önce, artan zorlukta 4-5 soruyla kalibrasyon yap.
  CV/sertifika/skill listesine güvenme — özellikle Nisan 2026'da kümelenmiş
  kısa online kurs sertifikaları (IBM/Google/Anthropic/Microsoft/AWS) sadece
  kelime hazinesi verir, hands-on derinlik vermez.

## Mimari
- `laravel-api/` — DDD'li çekirdek backend (Eloquent, Redis/Horizon)
- `python-ai/` — RAG + agent (asyncio, pgvector)
- `mcp-server/` — laravel-api'yi tool olarak agent'a açan köprü
- `go-ingestion/` — webhook event'lerini laravel-api'nin kuyruğuna basan servis

Her modül bağımsız yazılır, test edilir, demo edilir — hiçbiri diğerinin
bitmesini şart koşmaz. Birleşme yalnızca iki noktada olur:
1. `docker-compose` ile dördü aynı stack'te (RAG + Laravel + Postgres + Redis).
2. `mcp-server` gerçek bağlantıyı kurduğunda + `go-ingestion` gerçek event
   bastığında.

## Notlar
- Proje şu an boş — `/init` çalıştırıldığında büyük bir kod taraması bekleme,
  bu dosyada yazılanları koru, yalnızca kod tabanı oluştukça üzerine ekle.
- Kod tabanı büyüdükçe `/init` tekrar çalıştırılabilir.
