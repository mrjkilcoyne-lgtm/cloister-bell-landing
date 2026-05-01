---
realiser: rite-of-creation
created: 2026-05-01T00:00:00Z
parent_mandate: matt — "scaffold the slew, antigravity fills content"
status: scaffolded
---

# STATUS — Cloister Bell landing

**Phase:** scaffolded (binary builds; templates render; deploy step is stubbed).

## Real

- `main.go` compiles to a single binary. Routes: `/`, `/thesis`, `/hardware`, `/api/signup`, `/healthz`, 404.
- Templates embedded via `go:embed`. Static assets embedded.
- Distroless `Dockerfile`. Image target ~10–14 MB.
- `site.css` is bespoke, no Tailwind, palette + typescale per brief.
- Accessibility: skip-link, focus-visible, semantic HTML, `prefers-reduced-motion`.
- htmx-driven signup form posts to `/api/signup` and returns an HTML partial inline.

## Placeholder

- `static/htmx.min.js` is a **stub** — drop the real htmx 2.x minified file before launch:
  `curl -L https://unpkg.com/htmx.org@2.0.4/dist/htmx.min.js > static/htmx.min.js`
- All copy on `/`, `/thesis`, `/hardware` is sketch — clearly marked with `<!-- PLACEHOLDER -->` HTML comments.
- `/api/signup` echoes; no real subscriber backend wired.
- Deploy step in `.github/workflows/deploy.yml` is commented — sovereign chooses Fly.io vs Cloud Run, then add the secret and uncomment.
- Domain unknown (mandate says Matt names later).
- No OG image art.

## Gaps requiring Matt's hands

1. Pick the deploy target. Drop the relevant secret. Uncomment the deploy job.
2. Pick the domain. Drop a DNS record. Update `<meta property="og:*">` if needed.
3. Pick a subscriber backend (Buttondown / ConvertKit / Listmonk / self-hosted). Wire `/api/signup`.
4. Replace `static/htmx.min.js` stub with the real file.
5. Fill the placeholder copy on the three content pages — or hand off to a content realiser with the warning that bridge names, vocabulary sizes, and any user counts must come from Matt, not invention.

## Why this exists in this shape

The artefact's thesis is anti-SPA: latency, binary footprint, predictability.
A site selling that thesis cannot itself ship 2 MB of JS to render five
paragraphs. Hence htmx + Go. The architecture *is* the argument.
