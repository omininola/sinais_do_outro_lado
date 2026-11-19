const html = document.documentElement;
const themeSelect = document.querySelector("#themeSelect");

function getCookie(name) {
  const match = document.cookie.match(new RegExp("(^| )" + name + "=([^;]+)"));
  return match ? decodeURIComponent(match[2]) : null;
}

const defaultTheme = getCookie("theme") || "sinais";
html.setAttribute("data-theme", defaultTheme);
themeSelect.value = defaultTheme;

themeSelect.addEventListener("change", async () => {
  const theme = themeSelect.value;
  const days = 7;
  html.setAttribute("data-theme", theme);
  document.cookie = `theme=${theme}; path=/; max-age=${days * 24 * 60 * 60 * 1000}; SameSite=Lax`;
});
