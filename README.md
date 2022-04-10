

# Kyoto Start

Small, preconfigured project starter.
Includes:
- kyoto itself
- zen: set of commonly used tools
- i18n: internationalizing helper
- uikit: set of components for rapid development
- tailwindcss: an utility-first CSS framework

To start, simply clone this template with desired project name:

```bash
git clone --recursive https://github.com/kyoto-framework/start.git <project-name>
```

Then you'll need to build statics with:

```bash
(cd static; npm i; npm run build)
```

And now you can run your project:

```bash
go run .
```

From this moment, the only thing is left - override git origin with your repository:

```bash
git remote set-url origin <your-repository>
```
