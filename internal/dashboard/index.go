package dashboard

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Provenance Collector</title>
<link rel="icon" href="https://raw.githubusercontent.com/nebari-dev/nebari-design/main/symbol/favicon.ico">
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap" rel="stylesheet">
<style>
  :root {
    --bg: #0F1015;
    --surface: #16161f;
    --surface-2: #1e1e2e;
    --border: #2a2a3d;
    --text: #e2e0f0;
    --muted: #8a8aab;
    --faint: #55556a;
    --primary: #BA18DA;
    --primary-dim: rgba(186,24,218,0.12);
    --accent: #EAB54E;
    --green: #20AAA1;
    --yellow: #EAB54E;
    --red: #e5574b;
    --font: 'Poppins', system-ui, sans-serif;
  }
  * { margin: 0; padding: 0; box-sizing: border-box; }
  body { font-family: var(--font); background: var(--bg); color: var(--text); line-height: 1.6; -webkit-font-smoothing: antialiased; }

  /* Nav */
  .nav { position: sticky; top: 0; z-index: 100; background: rgba(15,15,20,0.85); backdrop-filter: blur(12px); border-bottom: 1px solid var(--border); }
  .nav-inner { max-width: 1260px; margin: 0 auto; padding: 0 24px; display: flex; align-items: center; justify-content: space-between; height: 56px; }
  .nav-brand { display: flex; align-items: center; gap: 10px; font-weight: 600; font-size: 15px; color: var(--text); }
  .nav-brand img { height: 24px; }
  .nav-brand .sep { color: var(--faint); font-weight: 300; }
  .nav-meta { font-size: 12px; color: var(--muted); display: flex; gap: 16px; }

  .container { max-width: 1260px; margin: 0 auto; padding: 24px; }

  /* Stats grid */
  .stats { display: grid; grid-template-columns: repeat(auto-fit, minmax(150px, 1fr)); gap: 12px; margin-bottom: 28px; }
  .stat { background: var(--surface); border: 1px solid var(--border); border-radius: 10px; padding: 18px; cursor: pointer; transition: all 0.2s; }
  .stat:hover { border-color: var(--primary); }
  .stat.active { border-color: var(--primary); box-shadow: 0 0 0 1px var(--primary), 0 0 20px rgba(186,24,218,0.15); }
  .stat .value { font-size: 26px; font-weight: 700; letter-spacing: -0.02em; }
  .stat .label { font-size: 11px; color: var(--muted); margin-top: 2px; text-transform: uppercase; letter-spacing: 0.05em; font-weight: 500; }
  .stat.green .value { color: var(--green); }
  .stat.yellow .value { color: var(--yellow); }
  .stat.red .value { color: var(--red); }

  /* Panels */
  .panel { background: var(--surface); border: 1px solid var(--border); border-radius: 10px; margin-bottom: 20px; overflow: hidden; }
  .panel-header { padding: 14px 20px; border-bottom: 1px solid var(--border); display: flex; justify-content: space-between; align-items: center; }
  .panel-header h2 { font-size: 13px; font-weight: 600; text-transform: uppercase; letter-spacing: 0.06em; color: var(--muted); }
  .panel-body { padding: 0; }

  /* Filters */
  .filters { padding: 10px 20px; border-bottom: 1px solid var(--border); display: flex; gap: 8px; flex-wrap: wrap; align-items: center; }
  .filters input[type="text"] {
    background: var(--bg); border: 1px solid var(--border); border-radius: 6px;
    padding: 6px 12px; color: var(--text); font-size: 12px; font-family: var(--font);
    min-width: 200px; outline: none; transition: border-color 0.2s;
  }
  .filters input[type="text"]::placeholder { color: var(--faint); }
  .filters input[type="text"]:focus { border-color: var(--primary); }
  .filters select {
    background: var(--bg); border: 1px solid var(--border); border-radius: 6px;
    padding: 6px 8px; color: var(--text); font-size: 12px; font-family: var(--font);
    outline: none; cursor: pointer;
  }
  .filters select:focus { border-color: var(--primary); }
  .filter-label { font-size: 11px; color: var(--faint); text-transform: uppercase; letter-spacing: 0.04em; font-weight: 500; }
  .filter-reset { background: none; border: 1px solid var(--border); border-radius: 6px; padding: 5px 10px; color: var(--muted); font-size: 11px; font-family: var(--font); cursor: pointer; margin-left: auto; transition: all 0.2s; }
  .filter-reset:hover { border-color: var(--primary); color: var(--text); }

  /* Tables */
  table { width: 100%; border-collapse: collapse; font-size: 13px; }
  th { text-align: left; padding: 10px 20px; color: var(--faint); font-weight: 500; font-size: 11px; text-transform: uppercase; letter-spacing: 0.05em; border-bottom: 1px solid var(--border); cursor: pointer; user-select: none; transition: color 0.2s; }
  th:hover { color: var(--muted); }
  th .sort-arrow { font-size: 9px; margin-left: 3px; }
  td { padding: 10px 20px; border-bottom: 1px solid var(--border); }
  tr:last-child td { border-bottom: none; }
  tr:hover { background: rgba(186,24,218,0.03); }

  /* Badges */
  .badge { display: inline-block; padding: 2px 8px; border-radius: 4px; font-size: 11px; font-weight: 500; letter-spacing: 0.02em; }
  .badge-green { background: rgba(32,170,161,0.12); color: var(--green); }
  .badge-yellow { background: rgba(234,181,78,0.12); color: var(--yellow); }
  .badge-red { background: rgba(229,87,75,0.12); color: var(--red); }
  .badge-muted { background: rgba(138,138,171,0.1); color: var(--faint); }
  .badge-primary { background: var(--primary-dim); color: #d46bec; }

  /* Timeline */
  .timeline { display: flex; gap: 10px; overflow-x: auto; padding: 14px 20px; }
  .timeline-item { min-width: 130px; padding: 10px 12px; border: 1px solid var(--border); border-radius: 8px; cursor: pointer; transition: all 0.2s; flex-shrink: 0; }
  .timeline-item:hover { border-color: var(--primary); }
  .timeline-item.active { border-color: var(--primary); background: var(--primary-dim); }
  .timeline-item .date { font-size: 12px; font-weight: 600; }
  .timeline-item .time { font-size: 11px; color: var(--muted); }
  .timeline-item .count { font-size: 11px; color: var(--faint); margin-top: 3px; }

  /* Misc */
  .empty { text-align: center; padding: 48px 20px; color: var(--muted); }
  .empty p { margin-top: 6px; font-size: 13px; }
  .loading { text-align: center; padding: 48px; color: var(--muted); font-size: 13px; }
  .mono { font-family: 'SF Mono', 'Fira Code', 'Consolas', monospace; font-size: 12px; }
  .text-muted { color: var(--muted); }
  .result-count { font-size: 11px; color: var(--faint); padding: 6px 20px; border-bottom: 1px solid var(--border); }
</style>
</head>
<body>
<nav class="nav">
  <div class="nav-inner">
    <div class="nav-brand">
      <img src="https://raw.githubusercontent.com/nebari-dev/nebari-design/main/symbol/nebari-favicon-dark.svg" alt="Nebari"
           onerror="this.style.display='none'">
      <span>Nebari</span>
      <span class="sep">/</span>
      <span>Provenance Collector</span>
    </div>
    <div class="nav-meta">
      <span id="cluster-name"></span>
      <span id="last-updated"></span>
    </div>
  </div>
</nav>

<div class="container">
  <div class="stats" id="stats">
    <div class="loading">Loading report data...</div>
  </div>

  <div class="panel">
    <div class="panel-header">
      <h2>Report Timeline</h2>
    </div>
    <div class="timeline" id="timeline"></div>
  </div>

  <div class="panel">
    <div class="panel-header">
      <h2>Container Images</h2>
      <span class="text-muted" style="font-size:12px" id="image-count"></span>
    </div>
    <div class="filters" id="image-filters"></div>
    <div id="image-result-count"></div>
    <div class="panel-body" id="images-table"></div>
  </div>

  <div class="panel">
    <div class="panel-header">
      <h2>Helm Releases</h2>
      <span class="text-muted" style="font-size:12px" id="helm-count"></span>
    </div>
    <div class="panel-body" id="helm-table"></div>
  </div>
</div>

<script>
let reports = [];
let currentReport = null;
let imageFilters = { search: '', namespace: '', signature: '', sbom: '', provenance: '', update: '' };
let statFilter = '';
let imageSortCol = '';
let imageSortAsc = true;

async function init() {
  try {
    const res = await fetch('/api/reports');
    reports = await res.json();
    if (!reports || reports.length === 0) { showEmpty(); return; }
    renderTimeline();
    await loadReport(reports[0].filename);
  } catch (e) {
    document.getElementById('stats').innerHTML = '<div class="empty"><p>Failed to load reports</p></div>';
  }
}

function showEmpty() {
  document.getElementById('stats').innerHTML = '';
  document.getElementById('timeline').innerHTML = '<div class="empty"><p>No reports yet. Run the collector to generate your first provenance report.</p></div>';
  document.getElementById('images-table').innerHTML = '';
  document.getElementById('helm-table').innerHTML = '';
}

function renderTimeline() {
  const el = document.getElementById('timeline');
  el.innerHTML = reports.map((r, i) => {
    const d = new Date(r.generatedAt);
    return '<div class="timeline-item' + (i === 0 ? ' active' : '') + '" onclick="selectReport(' + i + ')" id="tl-' + i + '">' +
      '<div class="date">' + d.toLocaleDateString() + '</div>' +
      '<div class="time">' + d.toLocaleTimeString() + '</div>' +
      '<div class="count">' + r.summary.totalImages + ' images</div>' +
    '</div>';
  }).join('');
}

async function selectReport(idx) {
  document.querySelectorAll('.timeline-item').forEach(el => el.classList.remove('active'));
  document.getElementById('tl-' + idx).classList.add('active');
  await loadReport(reports[idx].filename);
}

async function loadReport(filename) {
  const res = await fetch('/api/reports/' + filename);
  currentReport = await res.json();
  resetFilters();
  renderStats(currentReport);
  renderFilters(currentReport);
  renderImages(currentReport);
  renderHelm(currentReport);

  const d = new Date(currentReport.metadata.generatedAt);
  document.getElementById('last-updated').textContent = d.toLocaleString();
  document.getElementById('cluster-name').textContent = currentReport.metadata.clusterName || '';
}

function resetFilters() {
  imageFilters = { search: '', namespace: '', signature: '', sbom: '', provenance: '', update: '' };
  statFilter = '';
}

function renderStats(r) {
  const s = r.summary;
  const pctSigned = s.uniqueImages ? Math.round(s.signedImages / s.uniqueImages * 100) : 0;
  const pctSbom = s.uniqueImages ? Math.round(s.imagesWithSBOM / s.uniqueImages * 100) : 0;
  const pctProv = s.uniqueImages ? Math.round((s.imagesWithProvenance || 0) / s.uniqueImages * 100) : 0;

  document.getElementById('stats').innerHTML =
    statCard('all', s.uniqueImages, 'Unique Images', '') +
    statCard('signed', s.signedImages, 'Signed (' + pctSigned + '%)', pctSigned > 80 ? 'green' : pctSigned > 50 ? 'yellow' : 'red') +
    statCard('verified', s.verifiedImages, 'Verified', s.verifiedImages > 0 ? 'green' : '') +
    statCard('provenance', s.imagesWithProvenance || 0, 'SLSA Provenance (' + pctProv + '%)', pctProv > 0 ? 'green' : '') +
    statCard('sbom', s.imagesWithSBOM, 'With SBOM (' + pctSbom + '%)', pctSbom > 50 ? 'green' : pctSbom > 0 ? 'yellow' : '') +
    statCard('updates', s.imagesWithUpdates, 'Updates Available', s.imagesWithUpdates > 0 ? 'yellow' : 'green') +
    statCard('helm', s.totalHelmReleases, 'Helm Releases', '');
}

function statCard(id, value, label, cls) {
  const active = statFilter === id ? ' active' : '';
  return '<div class="stat ' + cls + active + '" onclick="toggleStatFilter(\'' + id + '\')" id="stat-' + id + '"><div class="value">' + value + '</div><div class="label">' + label + '</div></div>';
}

function toggleStatFilter(id) {
  statFilter = statFilter === id ? '' : id;
  document.querySelectorAll('.stat').forEach(el => el.classList.remove('active'));
  if (statFilter) document.getElementById('stat-' + statFilter).classList.add('active');
  imageFilters = { search: imageFilters.search, namespace: imageFilters.namespace, signature: '', sbom: '', provenance: '', update: '' };
  switch (statFilter) {
    case 'signed': imageFilters.signature = 'signed'; break;
    case 'verified': imageFilters.signature = 'verified'; break;
    case 'sbom': imageFilters.sbom = 'yes'; break;
    case 'provenance': imageFilters.provenance = 'yes'; break;
    case 'updates': imageFilters.update = 'yes'; break;
  }
  syncFilterUI();
  renderImages(currentReport);
}

function renderFilters(r) {
  const namespaces = [...new Set((r.images || []).map(img => img.namespace))].sort();
  document.getElementById('image-filters').innerHTML =
    '<input type="text" id="f-search" placeholder="Search image, workload..." oninput="onFilterChange()">' +
    '<span class="filter-label">Namespace</span>' +
    '<select id="f-namespace" onchange="onFilterChange()"><option value="">All</option>' +
      namespaces.map(ns => '<option value="' + esc(ns) + '">' + esc(ns) + '</option>').join('') + '</select>' +
    '<span class="filter-label">Signature</span>' +
    '<select id="f-signature" onchange="onFilterChange()"><option value="">All</option><option value="verified">Verified</option><option value="signed">Signed</option><option value="unsigned">Unsigned</option></select>' +
    '<span class="filter-label">SBOM</span>' +
    '<select id="f-sbom" onchange="onFilterChange()"><option value="">All</option><option value="yes">Has SBOM</option><option value="no">None</option></select>' +
    '<span class="filter-label">Provenance</span>' +
    '<select id="f-provenance" onchange="onFilterChange()"><option value="">All</option><option value="yes">Has SLSA</option><option value="no">None</option></select>' +
    '<span class="filter-label">Update</span>' +
    '<select id="f-update" onchange="onFilterChange()"><option value="">All</option><option value="yes">Available</option><option value="no">Current</option></select>' +
    '<button class="filter-reset" onclick="clearFilters()">Clear</button>';
}

function syncFilterUI() {
  const ids = ['search','namespace','signature','sbom','provenance','update'];
  ids.forEach(id => { const el = document.getElementById('f-' + id); if (el) el.value = imageFilters[id]; });
}

function onFilterChange() {
  imageFilters.search = (document.getElementById('f-search').value || '').toLowerCase();
  imageFilters.namespace = document.getElementById('f-namespace').value;
  imageFilters.signature = document.getElementById('f-signature').value;
  imageFilters.sbom = document.getElementById('f-sbom').value;
  imageFilters.provenance = document.getElementById('f-provenance').value;
  imageFilters.update = document.getElementById('f-update').value;
  statFilter = '';
  document.querySelectorAll('.stat').forEach(el => el.classList.remove('active'));
  renderImages(currentReport);
}

function clearFilters() {
  resetFilters(); syncFilterUI();
  document.querySelectorAll('.stat').forEach(el => el.classList.remove('active'));
  renderImages(currentReport);
}

function filterImages(images) {
  return images.filter(img => {
    if (imageFilters.search) {
      const hay = (img.image + ' ' + img.namespace + ' ' + img.workload.kind + '/' + img.workload.name).toLowerCase();
      if (!hay.includes(imageFilters.search)) return false;
    }
    if (imageFilters.namespace && img.namespace !== imageFilters.namespace) return false;
    if (imageFilters.signature) {
      const sig = img.signature;
      if (imageFilters.signature === 'verified' && !(sig && sig.verified)) return false;
      if (imageFilters.signature === 'signed' && !(sig && sig.signed)) return false;
      if (imageFilters.signature === 'unsigned' && sig && sig.signed) return false;
    }
    if (imageFilters.sbom) {
      const has = img.sbom && img.sbom.hasSBOM;
      if (imageFilters.sbom === 'yes' && !has) return false;
      if (imageFilters.sbom === 'no' && has) return false;
    }
    if (imageFilters.provenance) {
      const has = img.provenance && img.provenance.hasProvenance;
      if (imageFilters.provenance === 'yes' && !has) return false;
      if (imageFilters.provenance === 'no' && has) return false;
    }
    if (imageFilters.update) {
      const has = img.update && img.update.updateAvailable;
      if (imageFilters.update === 'yes' && !has) return false;
      if (imageFilters.update === 'no' && has) return false;
    }
    return true;
  });
}

function sortImages(images) {
  if (!imageSortCol) return images;
  const sorted = [...images];
  sorted.sort((a, b) => {
    let va, vb;
    switch (imageSortCol) {
      case 'image': va = a.image; vb = b.image; break;
      case 'namespace': va = a.namespace; vb = b.namespace; break;
      case 'workload': va = a.workload.kind + '/' + a.workload.name; vb = b.workload.kind + '/' + b.workload.name; break;
      case 'signature':
        va = a.signature ? (a.signature.verified ? 2 : a.signature.signed ? 1 : 0) : -1;
        vb = b.signature ? (b.signature.verified ? 2 : b.signature.signed ? 1 : 0) : -1;
        return imageSortAsc ? va - vb : vb - va;
      case 'provenance':
        va = a.provenance && a.provenance.hasProvenance ? 1 : 0;
        vb = b.provenance && b.provenance.hasProvenance ? 1 : 0;
        return imageSortAsc ? va - vb : vb - va;
      case 'sbom':
        va = a.sbom && a.sbom.hasSBOM ? 1 : 0;
        vb = b.sbom && b.sbom.hasSBOM ? 1 : 0;
        return imageSortAsc ? va - vb : vb - va;
      case 'update':
        va = a.update && a.update.updateAvailable ? 1 : 0;
        vb = b.update && b.update.updateAvailable ? 1 : 0;
        return imageSortAsc ? va - vb : vb - va;
      default: return 0;
    }
    if (typeof va === 'string') return imageSortAsc ? va.localeCompare(vb) : vb.localeCompare(va);
    return 0;
  });
  return sorted;
}

function onSortClick(col) {
  if (imageSortCol === col) { imageSortAsc = !imageSortAsc; } else { imageSortCol = col; imageSortAsc = true; }
  renderImages(currentReport);
}

function sortArrow(col) {
  if (imageSortCol !== col) return '';
  return '<span class="sort-arrow">' + (imageSortAsc ? ' &#9650;' : ' &#9660;') + '</span>';
}

function renderImages(r) {
  if (!r.images || r.images.length === 0) {
    document.getElementById('images-table').innerHTML = '<div class="empty"><p>No images discovered</p></div>';
    document.getElementById('image-count').textContent = '';
    document.getElementById('image-result-count').innerHTML = '';
    return;
  }

  const filtered = sortImages(filterImages(r.images));
  const total = r.images.length;
  document.getElementById('image-count').textContent = total + ' images';

  const hasFilters = Object.values(imageFilters).some(v => v);
  document.getElementById('image-result-count').innerHTML = hasFilters
    ? '<div class="result-count">Showing ' + filtered.length + ' of ' + total + '</div>' : '';

  if (filtered.length === 0) {
    document.getElementById('images-table').innerHTML = '<div class="empty"><p>No images match the current filters</p></div>';
    return;
  }

  let html = '<table><thead><tr>' +
    '<th onclick="onSortClick(\'image\')">Image' + sortArrow('image') + '</th>' +
    '<th onclick="onSortClick(\'namespace\')">Namespace' + sortArrow('namespace') + '</th>' +
    '<th onclick="onSortClick(\'workload\')">Workload' + sortArrow('workload') + '</th>' +
    '<th onclick="onSortClick(\'signature\')">Signature' + sortArrow('signature') + '</th>' +
    '<th onclick="onSortClick(\'provenance\')">Provenance' + sortArrow('provenance') + '</th>' +
    '<th onclick="onSortClick(\'sbom\')">SBOM' + sortArrow('sbom') + '</th>' +
    '<th onclick="onSortClick(\'update\')">Update' + sortArrow('update') + '</th>' +
    '</tr></thead><tbody>';

  for (const img of filtered) {
    const sig = img.signature
      ? (img.signature.verified ? badge('Verified', 'green') : img.signature.signed ? badge('Signed', 'yellow') : badge('Unsigned', 'red'))
      : badge('N/A', 'muted');
    const prov = img.provenance && img.provenance.hasProvenance
      ? badge('SLSA', 'primary') : badge('None', 'muted');
    const sbom = img.sbom && img.sbom.hasSBOM
      ? badge(img.sbom.format.toUpperCase(), 'green') : badge('None', 'muted');
    const update = img.update && img.update.updateAvailable
      ? badge(img.update.latestInMajor || img.update.newestAvailable, 'yellow') : badge('Current', 'green');
    const digest = img.digest ? '<span class="text-muted mono" style="font-size:11px">' + img.digest.substring(7, 19) + '</span>' : '';

    html += '<tr>' +
      '<td><span class="mono">' + esc(img.image) + '</span><br>' + digest + '</td>' +
      '<td style="font-size:12px">' + esc(img.namespace) + '</td>' +
      '<td style="font-size:12px">' + esc(img.workload.kind) + '/' + esc(img.workload.name) + '</td>' +
      '<td>' + sig + '</td>' +
      '<td>' + prov + '</td>' +
      '<td>' + sbom + '</td>' +
      '<td>' + update + '</td>' +
    '</tr>';
  }
  html += '</tbody></table>';
  document.getElementById('images-table').innerHTML = html;
}

function renderHelm(r) {
  if (!r.helmReleases || r.helmReleases.length === 0) {
    document.getElementById('helm-table').innerHTML = '<div class="empty"><p>No Helm releases discovered</p></div>';
    document.getElementById('helm-count').textContent = '';
    return;
  }
  document.getElementById('helm-count').textContent = r.helmReleases.length + ' releases';

  let html = '<table><thead><tr><th>Release</th><th>Namespace</th><th>Chart</th><th>Version</th><th>App Version</th><th>Status</th></tr></thead><tbody>';
  for (const hr of r.helmReleases) {
    const status = hr.status === 'deployed' ? badge('Deployed', 'green') : badge(hr.status, 'yellow');
    html += '<tr>' +
      '<td>' + esc(hr.releaseName) + '</td>' +
      '<td style="font-size:12px">' + esc(hr.namespace) + '</td>' +
      '<td class="mono">' + esc(hr.chart) + '</td>' +
      '<td class="mono">' + esc(hr.version) + '</td>' +
      '<td class="mono">' + esc(hr.appVersion) + '</td>' +
      '<td>' + status + '</td>' +
    '</tr>';
  }
  html += '</tbody></table>';
  document.getElementById('helm-table').innerHTML = html;
}

function badge(text, cls) { return '<span class="badge badge-' + cls + '">' + esc(text) + '</span>'; }
function esc(s) { const d = document.createElement('div'); d.textContent = s || ''; return d.innerHTML; }

init();
</script>
</body>
</html>
` + ""
