const { chromium } = require('@playwright/test');

const baseURL = process.env.WEB_URL || 'http://web_app-cnpj-validacao';

async function retry(fn, attempts = 30, delayMs = 1000) {
  let lastErr;
  for (let i = 0; i < attempts; i++) {
    try {
      return await fn();
    } catch (e) {
      lastErr = e;
      await new Promise(r => setTimeout(r, delayMs));
    }
  }
  throw lastErr;
}

(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage();

  await retry(async () => {
    const resp = await page.goto(baseURL, { waitUntil: 'domcontentloaded', timeout: 3000 });
    if (!resp || !resp.ok()) throw new Error('web not ready');
  });

  await page.waitForSelector('[data-testid="cnpj-input"]');
  await page.fill('[data-testid="cnpj-input"]', '04.252.011/0001-10');

  await page.waitForSelector('[data-testid="valid-value"]');
  const validText = await page.textContent('[data-testid="valid-value"]');
  const normalizedText = await page.textContent('[data-testid="normalized-value"]');

  if ((validText || '').trim() !== 'true') {
    throw new Error(`expected valid true, got: ${validText}`);
  }
  if ((normalizedText || '').trim() !== '04252011000110') {
    throw new Error(`expected normalized 04252011000110, got: ${normalizedText}`);
  }

  await browser.close();
  console.log('e2e ok');
})().catch((err) => {
  console.error(err);
  process.exit(1);
});
