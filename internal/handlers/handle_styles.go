package handlers

import "net/http"

func StylesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/css; charset=UTF-8")
	w.Write([]byte(styles))
}

const styles = `
body {
  font-family: Arial, Helvetica, sans-serif;

  --border-color: #ddd;
  --gray-100: #fafafa;
  --gray-500: #666;
  --gray-700: #333;
  --primary-100: #d2e7ff;
  --primary-200: #b1d3fa;
  --primary-400: #2e8fff;
  --primary-500: #0078ff;
  --primary-600: #0066db;
  --error-500: #f00;
  --error-100: #fee;
}

/*
  Josh's Custom CSS Reset
  https://www.joshwcomeau.com/css/custom-css-reset/
*/
*, *::before, *::after {
  box-sizing: border-box;
}
* {
  margin: 0;
}
body {
  line-height: 1.5;
  -webkit-font-smoothing: antialiased;
}
img, picture, video, canvas, svg {
  display: block;
  max-width: 100%;
}
input, button, textarea, select {
  font: inherit;
}
p, h1, h2, h3, h4, h5, h6 {
  overflow-wrap: break-word;
}
#root, #__next {
  isolation: isolate;
}

/*
  Common components
*/
button, .button {
  display: inline-block;
  border: 0;
  background-color: var(--primary-500);
  border-radius: 50px;
  color: #fff;
  padding: 2px 10px;
  cursor: pointer;
  text-decoration: none;
}
button:hover, .button:hover {
  background: var(--primary-400);
}

/*
  Custom components
*/
.error {
  background-color: var(--error-100);
  color: var(--error-500);
  text-align: center;
  padding: 1rem;
  display: none;
}
.error.visible {
  display: block;
}

#header {
  background-color: #fff;
  text-align: center;
  padding: 0.5rem 0 1.5rem;
}

#header h1 {
  font-size: 5rem;
}

.container {
  display: flex;
  flex-direction: column;
  gap: 4px;
  margin: 0 auto;
  max-width: 600px;
  padding: 20px;
}

.card {
  /* border: 1px solid var(--border-color); */
  border-radius: 6px;
  padding: 10px 16px;
  background-color: #fff;
}
.card > :first-child {
  margin-top: 0;
}
.card > :last-child {
  margin-bottom: 0;
}

.session-form {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.signin-form {
  display: flex;
  flex-direction: row;
  gap: 6px;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  padding: 10px 16px;
  background-color: #fff;
}

.signin-form input {
  flex: 1;
  border: 0;
}

.status-options {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 8px;
  margin: 10px 0;
}

.status-option {
  font-size: 2rem;
  width: 3rem;
  height: 3rem;
  padding: 0;
  background-color: #fff;
  border: 1px solid var(--border-color);
  border-radius: 3rem;
  text-align: center;
  box-shadow: 0 1px 4px #0001;
  cursor: pointer;
}

.status-option:hover {
  background-color: var(--primary-100);
  box-shadow: 0 0 0 1px var(--primary-400);
}

.status-option.selected {
  box-shadow: 0 0 0 1px var(--primary-500);
  background-color: var(--primary-100);
}

.status-option.selected:hover {
  background-color: var(--primary-200);
}

.status-line {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 10px;
  position: relative;
  margin-top: 15px;
}

.status-line:not(.no-line)::before {
  content: '';
  position: absolute;
  width: 2px;
  background-color: var(--border-color);
  left: 1.45rem;
  bottom: calc(100% + 2px);
  height: 15px;
}

.status-line .status {
  font-size: 2rem;
  background-color: #fff;
  width: 3rem;
  height: 3rem;
  border-radius: 1.5rem;
  text-align: center;
  border: 1px solid var(--border-color);
}

.status-line .desc {
  color: var(--gray-500);
}

.status-line .author {
  color: var(--gray-700);
  font-weight: 600;
  text-decoration: none;
}

.status-line .author:hover {
  text-decoration: underline;
}

.signup-cta {
  text-align: center;
  text-wrap: balance;
  margin-top: 1rem;
}
`
