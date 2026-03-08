import { chromium } from 'playwright';

const baseUrl = process.env.BASE_URL || 'http://mobile_app-cnpj-validacao/';

async function waitForServer(url, attempts = 90, delayMs = 500) {
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

(async () => {
  await waitForServer(baseUrl);

  const browser = await chromium.launch();
  const page = await browser.newPage();

  await page.goto(baseUrl, { waitUntil: 'networkidle' });

  const input = page.locator('#cnpj-input');
  await input.fill('04.252.011/0001-10');
  await page.waitForTimeout(200);

  const normalized = await page.locator('#result-normalized').innerText();
  const valid = await page.locator('#result-valid').innerText();

  if (!normalized.includes('04252011000110')) {
    throw new Error(`Expected normalized to contain 04252011000110, got: ${normalized}`);
  }
  if (!valid.includes('true')) {
    throw new Error(`Expected valid to contain true, got: ${valid}`);
  }

  await input.fill('11.111.111/1111-11');
  await page.waitForTimeout(200);
  const valid2 = await page.locator('#result-valid').innerText();
  if (!valid2.includes('false')) {
    throw new Error(`Expected valid to contain false, got: ${valid2}`);
  }

  await browser.close();
  console.log('e2e ok');
})().catch((err) => {
  console.error(err);
  process.exit(1);
});
