module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  darkMode: "class",
  theme: {
    extend: {
      backgroundColor: {
        app: "var(--color-bg-app)",
        counter: "var(--color-bg-counter)",
        report: "var(--color-bg-report)",
        "table-head": "var(--color-bg-table-head)",
      },
      divideColor: {
        default: "var(--color-border-divide)",
      },
      textColor: {
        "alert-inverse": "var(--color-alert-inverse)",
        "info-inverse": "var(--color-info-inverse)",
        "ok-inverse": "var(--color-ok-inverse)",
        default: "var(--color-text-default)",
        "table-head": "var(--color-table-head)",
      },
      colors: {
        background: "var(--color-background)",
        foreground: "var(--color-foreground)",
        "foreground-light": "var(--color-foreground-light)",
        "foreground-lighter": "var(--color-foreground-lighter)",
        "foreground-lightest": "var(--color-foreground-lightest)",
        alert: "var(--color-alert)",
        "alert-light": "var(--color-alert-light)",
        orange: "var(--color-orange)",
        yellow: "var(--color-yellow)",
        ok: "var(--color-ok)",
        info: "var(--color-info)",
        tbd: "var(--color-tbd)",
        link: "var(--color-link)",
        "table-border": "var(--color-table-border)",
        "table-divide": "var(--color-table-divide)",
        "table-head": "var(--color-table-head)",
        "steampipe-black": "#181717",
        "black-scale-1": "var(--color-black-scale-1)",
        "black-scale-2": "var(--color-black-scale-2)",
        "black-scale-3": "var(--color-black-scale-3)",
        "black-scale-4": "var(--color-black-scale-4)",
        "black-scale-5": "var(--color-black-scale-5)",
        "black-scale-6": "var(--color-black-scale-6)",
        "black-scale-7": "var(--color-black-scale-7)",
        "black-scale-8": "var(--color-black-scale-8)",
      },
      fontSize: {
        xxs: ".65rem",
      },
      maxHeight: {
        "1/2-screen": "50vh",
      },
      screens: {
        print: { raw: "print" },
      },
      typography: (theme) => ({
        DEFAULT: {
          css: {
            color: theme("colors.foreground"),
            a: {
              color: theme("colors.link"),
              "&:hover": {
                color: theme("colors.link"),
              },
            },
            code: { color: theme("colors.foreground") },
            "a code": { color: theme("colors.foreground") },
            h1: { color: theme("colors.foreground") },
            h2: { color: theme("colors.foreground") },
            h3: { color: theme("colors.foreground") },
            h4: { color: theme("colors.foreground") },
            h5: { color: theme("colors.foreground") },
            h6: { color: theme("colors.foreground") },
            strong: { color: theme("colors.foreground") },
            "thead tr th": {
              color: theme("colors.table-head"),
            },
            "tbody tr": { borderBottomColor: theme("colors.table-divide") },
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};