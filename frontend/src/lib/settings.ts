import { SettingsService } from "../../bindings/github.com/sillsdev/appbuilder-support-data-collector";
import { locales, setLocale, type Locale } from "./paraglide/runtime.js";

/** daisyUI theme for each -theme value (see app.css). */
const themes: Record<string, string> = {
  light: "corporate",
  dark: "dim",
};

/**
 * The supported locale that best matches a language ID such as "fr", "pt_BR"
 * or "zh-cn": an exact match ignoring case, else the first locale with the same
 * primary language.
 */
function matchLocale(language: string): Locale | undefined {
  const wanted = language.trim().replace(/_/g, "-").toLowerCase();
  if (!wanted) {
    return undefined;
  }
  const primary = (tag: string) => tag.split("-")[0].toLowerCase();
  return (
    locales.find((locale) => locale.toLowerCase() === wanted) ??
    locales.find((locale) => primary(locale) === primary(wanted))
  );
}

/** Applies the theme and language passed on the command line. */
export async function applyStartupSettings(): Promise<void> {
  const settings = await SettingsService.GetSettings();

  const theme = themes[settings.theme.trim().toLowerCase()];
  if (theme) {
    document.documentElement.dataset.theme = theme;
  }

  const locale = matchLocale(settings.language);
  if (locale) {
    setLocale(locale);
  }
}
