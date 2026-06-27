# 12 Haftalık Birleşik Çalışma Planı (v5 — Go & Agentic AI)
### Agentic Customer Data Platform (CDP) Stack

Bu plan, modern backend pratikleri (Go, DDD, Clean Architecture) ile güncel AI trendlerini (RAG, AI Agents, MCP) bir araya getirerek production-grade ve CV-ready bir monorepo sistemi inşa etmeyi amaçlar.

---

## Tek Sistem: Ne İnşa Ediyoruz?

CDP (Customer Data Platform) / MarTech temalı, **müşteri, segment ve davranışsal event** verilerini işleyen akıllı bir mikroservis ekosistemi:

| Parça | Rol | Teknoloji |
|---|---|---|
| **`go-core-api`** | DDD'li çekirdek backend — müşteri/segment yönetimi ve event ingestion | Go (net/http veya Gin), Gorm, PostgreSQL, Redis |
| **`ai-agent-service`** | RAG (bilgi bankası) + ReAct Agent (Go API'sini dinamik sorgulama) | Python, FastAPI, pgvector, LangGraph / LangChain |
| **`mcp-gateway`** | Go API'sini AI Agent'a güvenli "Tool" olarak sunan köprü | Python/Node, Model Context Protocol (MCP) SDK |

**Production ve DevOps Vizyonu:**
- Her servis en başından itibaren Dockerize edilerek izole çalıştırılacak.
- AWS (VPC, EKS/K8s, IAM Instance Profiles) üzerinde zero-downtime (rolling update) deployment altyapısı kurulacak.

---

## FAZ 1 — Go Core Backend & Containerization (Hafta 1-3)

**🟦 Go Core Backend (DDD & Clean Architecture)**
- **Hafta 1:** Go projesinin standart layout (`cmd`, `internal`, `pkg`) ile kurulması. DDD prensiplerine göre katmanlama (Domain, Application, Infrastructure, Interfaces). Temel Customer & Segment CRUD modelleri ve veritabanı (PostgreSQL) entegrasyonu.
- **Hafta 2:** Event Ingestion Endpoint'leri. Yüksek eşzamanlılık (High Concurrency) için Go'nun native `goroutine` ve `channel` yapılarının kullanılması. Event validation ve buffering mekanizmaları.
- **Hafta 3:** Servisin `Dockerfile` (multi-stage build) ve `docker-compose` ortamına taşınması. Unit testler, mock kütüphaneleri ile Go test kapsamının artırılması.

---

## FAZ 2 — RAG & Vector DB + AI Agent Service (Hafta 4-7)

**🟩 AI Agent Service (Python & FastAPI)**
- **Hafta 4:** Python FastAPI tabanlı asenkron yapının kurulması. Chunking ve Embedding (Hugging Face / OpenAI) mekanizmaları.
- **Hafta 5:** Vector Database olarak PostgreSQL + `pgvector` kullanımı. Baseline RAG pipeline geliştirilmesi.
- **Hafta 6:** Hybrid Search (Semantic + BM25) ve Reranking (Cross-Encoder) entegrasyonu. Doküman yükleme ve indeksleme akışı.
- **Hafta 7:** AI Agent Core. LangGraph/LangChain ile ReAct loop tasarlanması ve bellek yönetimi (Chat History). Python servisinin Dockerize edilmesi.

---

## FAZ 3 — MCP Gateway & Entegrasyon (Hafta 8-9)

**🟪 Model Context Protocol (MCP) Bridge**
- **Hafta 8:** MCP standartlarının uygulanması. Go API'sinin endpoint'lerini (`go-core-api`) AI Agent'a `tools` olarak sunacak bir MCP sunucusunun yazılması.
- **Hafta 9:** **Entegrasyon Noktası:** Go API + Postgres + Python AI Agent + MCP sunucusunun tek bir `docker-compose` dosyası altında ayağa kaldırılması. Agent'ın Go API üzerinden canlı veri çekerek kullanıcı sorularını yanıtlamasının doğrulanması.

---

## FAZ 4 — AWS DevOps & Kubernetes (Hafta 10-11)

**🛡️ Cloud & Infrastructure**
- **Hafta 10:** AWS temelleri. VPC tasarımı (Public/Private subnets, NAT Gateway yerine VPC Endpoints ile maliyet kontrolü). EC2 IAM instance profiles. Kubernetes (EKS veya hafif K8s) temelleri (Pod, Deployment, Service, ConfigMap, Secret).
- **Hafta 11:** K8s Manifest'lerinin yazılması ve rolling update (zero-downtime) stratejisinin simülasyonu. Başarısız deploy durumunda `rollout undo` ile otomatik kurtarma senaryoları.

---

## FAZ 5 — Toparlama & Vitrin (Hafta 12)

- Monoreponun dökümantasyonunun tamamlanması. Servis bazlı `README.md` dosyalarının oluşturulması.
- Mimariyi gösteren Mermaid diyagramlarının güncellenmesi.
- Portfolyoyu LinkedIn ve CV'ye ekleme aşaması.
