# 12 Haftalık Birleşik Çalışma Planı (v4 — Final)
### Insider'a Hazırlık + AI Engineering — Tek Sistem Üzerinden

> **Bu sürüm nasıl bu hale geldi:**
> - **v1 → v2:** Laravel'in "derinleştirme" değil, sıfırdan öğrenme olduğu netleşti (CV'de PHP/Laravel hiç yok).
> - **v2 → v3:** Docker/K8s/AWS/MCP için "zaten biliniyor" varsayımı **kalibrasyon testleriyle** çöktü — gerçek seviyeye göre yeniden zamanlandı.
> - **v3 → v4:** İki ayrı track, **tek bir sistem**de birleştirildi (gevşek bağlı modüller + ortak domain + 2 entegrasyon noktası). Go'nun Insider'da PHP'ye **eşdeğer ağırlıkta bir backend dili** olduğu ilan verisiyle doğrulandı — Laravel/DDD'nin ilan dilindeki ("modüler monolit", "DDD", "PHPStan/Rector/Pint") birebir eşleşmesi de doğrulandı.
>
> **Tempo:** ~12–15 saat/hafta. **%40 Insider + %60 AI çekirdeği.**

---

## Tek sistem: ne inşa ediyoruz

CDP/martech temalı, **müşteri/segment/etkinlik** domain'i üzerine kurulu bir sistem:

| Parça | Rol | Teknoloji |
|---|---|---|
| **Laravel API** | DDD'li çekirdek backend — segment/müşteri yönetimi | PHP 8.x, Eloquent, Redis/Horizon |
| **Go servisi** | Event ingestion — webhook'ları alır, Laravel'in kuyruğuna basar | Go, net/http |
| **Python AI servisi** | RAG (statik bilgi) + Agent (canlı veri sorgusu) | Python, asyncio, pgvector |
| **MCP server** | Laravel API'yi agent'a tool olarak açan köprü | Python/Node, MCP SDK |

**Gevşek bağlılık ilkesi:** her parça bağımsız yazılır, test edilir, tek başına demo edilebilir. Birleşme yalnızca iki noktada olur:
- **Entegrasyon Noktası 1** (Faz 2 sonu): hepsi aynı `docker-compose` stack'inde yan yana çalışır.
- **Entegrasyon Noktası 2** (Faz 3 sonu): MCP server gerçek bağlantıyı kurar, Go servisi gerçek event besler.

Aradaki hiçbir hafta bir öncekinin bitmesini şart koşmaz.

---

## FAZ 1 — Temel (Hafta 1–3)

**🟦 Insider — Laravel'e giriş (sıfırdan, ilan diliyle doğrulandı)**
- **Hafta 1:** Routing, controller, Eloquent ORM, migration, Artisan. Basit bir kaynak (müşteri/segment) için CRUD REST API.
- **Hafta 2:** PHP 8.x modern pratikler + tooling. Strict types, readonly class, enum. PHPStan (max level) + Rector + Pint, Hafta 1 API'sine uygula.
- **Hafta 3:** DDD katmanlama. Entity / Repository / Manager / DTO ayrımıyla yeniden yapılandır — bu artık "ilan diliyle birebir örtüşen" gerçek bir backend pattern'i.

**🟩 AI çekirdeği**
- **Hafta 1:** Python async + type hints. *Not: Node.js event loop / Dart async-await zaten biliniyor, sözdizimi geçişi olacak — ama gerçek ilerleme henüz Claude Code'da başlamadı.*
- **Hafta 2:** LLM API + streaming CLI chatbot. *Planlama sohbetinde taslak (`chatbot.py`, `sync_vs_async.py`, `requirements.txt`) hazırlandı — referans olarak kullanılabilir, ama resmi ilerleme sayılmıyor.*
- **Hafta 3:** Tool calling — Pydantic şema, mock DB tool, retry. *(MCP kalibrasyonunda görülen `stop_reason`/`tool_result` karışıklığı tam burada netleşecek.)*

> **✅ Çıktı:** DDD'li Laravel API (sistemin backend çekirdeği) + streaming & tool-calling yapan CLI chatbot.

---

## FAZ 2 — RAG & Vector DB + Insider Altyapı (Hafta 4–7)

**🟩 AI çekirdeği**
- **Hafta 4:** Embeddings + chunking (semantik sınırla).
- **Hafta 5:** pgvector ile vector DB.
- **Hafta 6:** Framework'süz baseline RAG pipeline.
- **Hafta 7:** Hybrid search (BM25 + dense) + reranker.

**🟦 Insider — AWS hedefli, Docker tam temelden**
- **Hafta 4:** Redis + Laravel Horizon — job/queue akışı.
- **Hafta 5:** **AWS — hedefli.** S3/Lambda/Elastic IP'yi tekrar etme (kalibrasyonda sağlam çıktı). Odak: **VPC** (public/private subnet, NAT Gateway) + **EC2 IAM instance profile** (kalibrasyonda net boşluk). Hands-on: private subnet'te bir EC2, instance profile ile S3'e erişsin.
- **Hafta 6:** **Docker — tam temelden.** Image/layer, multi-stage build, networking, volume, restart-loop debug. Sıkıştırma yok (kalibrasyon: geniş boşluk).
- **Hafta 7:** Docker devamı + **Entegrasyon Noktası 1**: RAG + Laravel + Postgres + Redis aynı `docker-compose` stack'inde.

> **✅ Çıktı:** VPC + instance profile ile çalışan bir AWS deploy'u + Docker fundamentals sağlam, compose'da **gerçekten birleşmiş** RAG+Laravel stack'i.

---

## FAZ 3 — Agent, MCP & Üretim (Hafta 8–11)

**🟩 AI çekirdeği**
- **Hafta 8:** Agent mimarisi — ReAct loop.
- **Hafta 9:** **MCP — önce mekanizma, sonra inşa.** Önce: client-host-server mimarisi, tools/resources/prompts üçlüsü (kalibrasyonda bilinmiyordu), gerçek tool_use akışı (MCP değil, çağıran uygulama tool'u çalıştırır). Sonra: gerçek bir MCP server yaz — bu, **Entegrasyon Noktası 2**'nin başlangıcı (Laravel API'yi tool olarak açar).
- **Hafta 10:** RAGAS evaluation (faithfulness / answer relevance / context relevance).
- **Hafta 11:** Observability (LangSmith/Phoenix) + kısa not: **prompt caching** = context/prefix cache (tekrar eden system prompt/tool tanımını yeniden işletmeme), semantic cache (benzer sorulara hazır cevap) **değil** — kalibrasyonda bu ikisi karışmıştı.

**🟦 Insider — K8s tam temelden, Go eşdeğer ciddiyette**
- **Hafta 8:** Kubernetes temelleri — Pod/Service/Deployment/ConfigMap/Secret, basit bir manifest + `kubectl apply`.
- **Hafta 9:** Faz 2'deki compose stack'ini K8s'e taşı. Rolling update + `kubectl rollout undo` (kalibrasyonda hiç bilinmiyordu).
- **Hafta 10:** Go'ya giriş — sözdizimi, goroutine'ler. *(Not: ilan verisi Go'yu "ilgi göstergesi" değil, PHP'ye **eşdeğer** bir backend dili olarak gösteriyor — buna göre ele al.)*
- **Hafta 11:** Go ile gerçek bir HTTP servisi — webhook event'lerini Laravel'in kuyruğuna basan ingestion servisi. **Entegrasyon Noktası 2 tamamlanır.** Hata yönetimi ve testi Laravel API ile aynı kalitede olsun — "giriş seviyesi ekstra" değil, gerçek bir parça.

> **✅ Çıktı:** Tools/resources/prompts'u doğru kullanan, kendi MCP server'ına bağlı, RAGAS'la ölçülmüş bir agent. K8s'te deploy edilmiş, rolling update'i test edilmiş bir stack. Laravel'e gerçek event basan bir Go servisi.

---

## FAZ 4 — Toparlama & Vitrin (Hafta 12)

- Sistemin parçalarını (Laravel API, Python AI servisi, MCP server, Go servisi) GitHub'da README'lerle topla — **tek bir mimari anlatımı**, dört dağınık demo değil.
- Kafka + ClickHouse'u **kavramsal** çalış — kod yazma, ama "event streaming + kolon-tabanlı analitik DB, CDP'de yeri ne" diyebil.
- LinkedIn/CV güncelle: "Laravel" satırının artık gerçek bir karşılığı var.
- Insider ilanlarını tara (özellikle "Backend Software Developer (PHP/Laravel)" ve Go/Node.js geçen ilanlar), en yakın eşleşen role başvur.

---

## Hızlı Referans (final)

| Sıra | Konu | Durum | Kanıt |
|---|---|---|---|
| 1 | Laravel temelleri → DDD | Sıfırdan | İlan diliyle birebir doğrulandı |
| 2 | Go | Sıfırdan, **eşdeğer ciddiyet** | İlanlarda "PHP ve Go/Node.js" — eşdeğer dil |
| 3 | Python async + LLM API | Henüz başlanmadı | Node.js/Dart transferi var, sözdizimi yeni |
| 4 | RAG + Vector DB | Planlanan | — |
| 5 | Docker fundamentals | Sıfırdan | Kalibrasyon testiyle doğrulandı |
| 6 | Kubernetes fundamentals | Sıfırdan, 2 hafta | Kalibrasyon testiyle doğrulandı |
| 7 | AWS — VPC + instance profile | Hedefli boşluk | Kalibrasyon testiyle doğrulandı, gerisi sağlam |
| 8 | Tool use API mekaniği | Hafta 3'te netleşir | Kalibrasyon testiyle doğrulandı |
| 9 | MCP primitives + server | Terim biliniyor, mekanizma yeni | Kalibrasyon testiyle doğrulandı |
| 10 | Prompt caching vs semantic cache | Kavram düzeltmesi | Kalibrasyon testiyle doğrulandı |

---

### Not
Hiçbir konu "zaten biliyor" diye atlanmadı — her varsayım ya CV/ilan kanıtıyla ya da
kalibrasyon testiyle doğrulandı. Yeni bir konuya geçerken aynı yöntem (artan
zorlukta 4-5 soru) kullanılabilir. **İlerleme: plan henüz uygulanmaya başlanmadı —
bu sohbet tamamen stratejik planlama oturumuydu.** Hafta 1'den itibaren resmi
ilerleme Claude Code'da, gerçek bir proje dizininde başlayacak.
