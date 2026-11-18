const html = document.documentElement;
const themeSelect = document.querySelector("#themeSelect");

themeSelect.addEventListener("change", () => {
  theme = themeSelect.value;
  html.setAttribute("data-theme", theme);
});
