
# Setup
## Install Hugo
https://gohugo.io/installation/macos/

```bash
brew install hugo
```

Verify install

```bash
hugo version
```

## Create a new site[#](https://jpanther.github.io/congo/docs/installation/#create-a-new-site)

Run the command `hugo new site mywebsite` to create a new Hugo site in a directory named `mywebsite`.

Install using git[#](https://jpanther.github.io/congo/docs/installation/#install-using-git)

For this method you’ll need to ensure you have **Git** installed on your local machine.

Change into the directory for your Hugo website (that you created above), initialise a new `git` repository and add Congo as a submodule.

```bash
cd mywebsite
git init
git submodule add -b stable https://github.com/jpanther/congo.git themes/congo
```

Then continue to [set up the theme configuration files](https://jpanther.github.io/congo/docs/installation/#set-up-theme-configuration-files).

## Add content

```bash
hugo new content content/posts/my-first-post.md
```


## Build and serve content

```text
hugo server --buildDrafts
```


## Publish the project

In this step you will _publish_ your project, but you will not _deploy_ it.

When you publish your project, Hugo renders all build artifacts to the `public` directory in the root of your project. This includes the HTML files for every site, along with assets such as images, CSS, and JavaScript. The command is simple.

```text
hugo
```


## Visual Studio Code plugins

[Front Matter](https://marketplace.visualstudio.com/items?itemName=eliostruyf.vscode-front-matter)

This extension maintains article metadata such as creation date, modified date, slug, title, SEO check, and more

[Hugo Language and Syntax Support](https://marketplace.visualstudio.com/items?itemName=budparr.language-hugo-vscode)

This extension provides syntax highlighting and snippets. The source code is available on its [GitHub repository](https://github.com/budparr/language-hugo-vscode).

---

## Quick References

### Native Hugo

Emojis on Markdown
https://gohugo.io/quick-reference/emojis/

Code Highlighting Styles
https://gohugo.io/quick-reference/syntax-highlighting-styles/

Hugo Shortcodes
https://gohugo.io/shortcodes/

Diagrams 
https://gohugo.io/content-management/diagrams/

Front matter
https://gohugo.io/content-management/front-matter/

### Congo Specific

Icon reference
https://jpanther.github.io/congo/samples/icons/

--- 

# Next steps

Add comments section to the blog
https://gohugo.io/content-management/comments/
https://utteranc.es/
