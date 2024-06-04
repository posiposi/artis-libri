import { Linter } from "eslint";

/** @type {Linter.Config} */
const config = {
  env: {
    browser: true,
    es2021: true
  },
  extends: [
    "eslint:recommended",
    "plugin:react/recommended",
  ],
  parserOptions: {
    ecmaVersion: 12,
    sourceType: "module"
  },
  plugins: [
    "react"
  ],
  settings: {
    react: {
      version: "detect"
    }
  },
  rules: {
    "react/react-in-jsx-scope": "off"
  }
};

export default config;
