const puppeteer = require('puppeteer');
const fs = require('fs');
const path = require('path');

async function generatePDF(htmlTemplatePath, outputPdfPath) {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    const htmlContent = fs.readFileSync(htmlTemplatePath, 'utf8');
    await page.setContent(htmlContent, { waitUntil: 'networkidle0' });
    await page.pdf({ path: outputPdfPath, format: 'A4', printBackground: true });
    await browser.close();
}

const htmlTemplatePath = path.join(__dirname, './invoice.html');
const outputPdfPath = path.join(__dirname, 'invoice.pdf');

generatePDF(htmlTemplatePath, outputPdfPath)
    .then(() => console.log('PDF generated successfully'))
    .catch(err => console.error('Error generating PDF:', err));
