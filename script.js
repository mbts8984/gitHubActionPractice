const CONFIG_FILE = './variables-to-check.yaml'
const yaml = require('js-yaml')
const fs = require('fs')


module.exports = async ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")

    try {
        const doc = yaml.load(fs.readFileSync(CONFIG_FILE, 'utf-8'))
        console.log('Docssss ', doc)
    } catch (error) {
        console.log('error getting yaml', error)
    }

    const matchesPattern = (pattern, text) => {
        const regex = new RegExp(pattern);
        return !!text.match(regex)
      };

    const diff_url = context.payload.pull_request.diff_url
    const files = await github.request(diff_url)
    
    const file = files.data
    const config = await context.config

    // console.log("CONFIG HERE: ", config)

    let position = 0
    file.split("\n").forEach((line) => {
        // console.log('made it inloop', line)
        if (line.startsWith("+")){
            const addedLine = line.slice(1);
            // console.log('++++ ', addedLine)
        }
    })

    try {
       await github.issues.createComment({
            issue_number: context.issue.number,
            owner: context.repo.owner,
            repo: context.repo.repo,
            body: ' 👋👋👋👋👋 test'
          })
    } catch (error) {
        console.log("error", error)
    }
  }