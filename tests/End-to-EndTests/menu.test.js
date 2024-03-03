const { test, expect } = require('@playwright/test');

test('submit support request form', async ({ page }) => {
  await page.goto('http://localhost:3000/support/');
  
  await page.fill('#fullName', 'John Doe');
  await page.fill('#email', 'johndoe@example.com');
  await page.fill('#subject', 'Problem with product A');
  await page.fill('#message', 'Hello, I need help with my order. I want something just like this but it does not work'
  );
  
  await page.click('button[type="submit"]');
  
  const confirmationAlert = page.locator('#confirmationAlert');
  await expect(confirmationAlert).toBeVisible();

  await expect(confirmationAlert).toContainText("Your request has been submitted successfully. We'll get back to you soon.");
});
