---
realiser: rite-of-creation
created: 2026-05-01T00:00:00Z
parent_mandate: matt — "scaffold the slew, antigravity fills content"
status: scaffolded
---

# Cloister Bell — landing site

Single-binary htmx + Go landing page for the Cloister Bell thought-decoder
artefact. The architecture *is* the thesis: server-rendered HTML, no SPA,
no bundler, no framework drift. One binary, embedded templates and static
assets, ~10 MB image, sub-millisecond page renders.

## Stack

- **Go 1.23** with `embed` for templates and static assets
- **htmx 2.x** for the small interactive bits (drop the real file into `static/htmx.min.js` — current file is a placeholder stub)
- Hand-rolled CSS, no Tailwind, no JS bundler
- Distroless container, ~12 MB final image

## Run locally

```bash
go run ./           # http://localhost:8080
go build ./         # produces ./cloister-bell binary
./cloister-bell -addr :9000
```

## Container

```bash
docker build -t cloister-bell-landing .
docker run --rm -p 8080:8080 cloister-bell-landing
```

## Deploy

The GitHub Actions workflow currently **builds only**. Two ratified deploy
targets are commented in `.github/workflows/deploy.yml`:

- **Fly.io** — drop `FLY_API_TOKEN` and a `fly.toml`, uncomment the deploy step.
- **Cloud Run** — drop `GCP_SA_KEY`, region/project, uncomment the deploy step.

Sovereign to choose; this scaffold supports either.

## Why htmx + Go and not SPA

The artefact this site fronts is a real-time biosignal decoder. Latency,
predictability and binary footprint matter. A 2 MB JS bundle to render
five paragraphs of marketing copy would be dishonest about what the
underlying product is. So: one binary. Server-rendered. htmx for the
two interactive bits.

See `STATUS.md` for what is real and what is placeholder.
