import { useState } from "react";
import { themeColors } from "../utils/themeColors";

export const useTheme = (name) => {
  const [theme, setTheme] = useState("light");

  const toggleTheme = () => {
    let name = theme === "light" ? "dark" : "light"

    document.body.style.setProperty("--primary", themeColors[name].primary);
    document.body.style.setProperty(
      "--primary-variant",
      themeColors[name].primaryVariant
    );
    document.body.style.setProperty("--secondary", themeColors[name].secondary);
    document.body.style.setProperty(
      "--secondary-variant",
      themeColors[name].secondaryVariant
    );
    document.body.style.setProperty("--background", themeColors[name].background);
    document.body.style.setProperty("--surface", themeColors[name].surface);
    document.body.style.setProperty("--error", themeColors[name].error);
    document.body.style.setProperty("--on-primary", themeColors[name].onPrimary);
    document.body.style.setProperty(
      "--on-secondary",
      themeColors[name].onSecondary
    );
    document.body.style.setProperty(
      "--on-background",
      themeColors[name].onBackground
    );
    document.body.style.setProperty("--on-surface", themeColors[name].onSurface);
    document.body.style.setProperty("--on-error", themeColors[name].onError);

    setTheme(name)
  }

  return [theme, toggleTheme];
};
