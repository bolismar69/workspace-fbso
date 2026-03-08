import { chromium } from 'playwright';

const baseUrl = process.env.BASE_URL || 'http://mobile_app-cnpj-validacao/';

async function waitForServer(url, attempts = 120, delayMs = 500) {
  for (let i = 0; i < attempts; i++) {
    try {
      const res = await fetch(url, { redirect: 'follow' });
      if (res.ok) return;
    } catch {
      // ignore
    }
    await new Promise((r) => setTimeout(r, delayMs));
  }
  throw new Error(`Server not ready at ${url}`);
}

async function readE2eData(page) {
  const json = await page.evaluate(() => globalThis.__e2eResultsJson ?? null);
  if (typeof json !== 'string') return null;
  try {
    return JSON.parse(json);
  } catch {
    return null;
  }
}

async function waitForE2eState(page, predicate, timeoutMs = 60000, stepMs = 100) {
  const deadline = Date.now() + timeoutMs;
  let last = null;
  while (Date.now() < deadline) {
    // eslint-disable-next-line no-await-in-loop
    const data = await readE2eData(page);
    if (data) {
      last = data;
      if (predicate(data)) return data;
    }
    // eslint-disable-next-line no-await-in-loop
    await page.waitForTimeout(stepMs);
  }
  throw new Error(`Timed out waiting for E2E state. Last seen: ${JSON.stringify(last)}`);
}

async function getCnpjInput(page) {
  // HTML renderer tends to expose a real input/textarea with placeholder.
  const byPlaceholder = page.locator(
    'textarea[placeholder="99.999.999/9999-99"], input[placeholder="99.999.999/9999-99"]'
  );
  if ((await byPlaceholder.count()) > 0) {
    return byPlaceholder.first();
  }

  // Fallback: Flutter's text editing host.
  const hostInput = page.locator('flt-text-editing-host textarea, flt-text-editing-host input').first();
  return hostInput;
}

(async () => {
  await waitForServer(baseUrl);

  const browser = await chromium.launch();
  const context = await browser.newContext({
    viewport: { width: 1280, height: 720 },
    locale: 'pt-BR',
  });
  const page = await context.newPage();

  page.on('console', (msg) => {
    if (msg.type() === 'error') {
      console.error(`[browser:console] ${msg.text()}`);
    }
  });
  page.on('pageerror', (err) => {
    console.error('[browser:pageerror]', err);
  });

  await page.goto(baseUrl, { waitUntil: 'networkidle' });

  // Wait until Flutter has written initial state.
  await waitForE2eState(page, (data) => data && Object.prototype.hasOwnProperty.call(data, 'valid'));

  const input = await getCnpjInput(page);
  // If we're on a renderer where the real <input>/<textarea> is created lazily,
  // tabbing into the first focusable control is the most reliable way to force it.
  await page.keyboard.press('Tab');
  await input.waitFor({ state: 'attached', timeout: 30000 });
  await input.click({ timeout: 30000 });
  await input.fill('04.252.011/0001-10');

  await waitForE2eState(page, (data) => data.normalized === '04252011000110');
  await waitForE2eState(page, (data) => data.valid === true);

  await input.click();
  await input.fill('11.111.111/1111-11');
  await waitForE2eState(page, (data) => data.valid === false);

  await browser.close();
  console.log('e2e ok');
})().catch((err) => {
  console.error(err);
  process.exit(1);
});
