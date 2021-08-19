const CONFIG_FILE = './variables-to-check.yaml'
const yaml = require('js-yaml')
const fs = require('fs')


module.exports = async ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")

    try {
        const config = yaml.load(fs.readFileSync(CONFIG_FILE, 'utf-8'))
        console.log('Docssss ', config)
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
    // const config = await context.config

    let position = 0
    file.split("\n").forEach((line) => {
        // console.log('made it inloop', line)
        if (line.startsWith("+")){
            const addedLine = line.slice(1);
            const matches = config.matches;

            matches.forEach((match) => {
                const commentText = "found a match here "

                if (matchesPattern(match.regex, addedLine)){
                    console.log("Commenting " + commentText + "on line " + position)
                    try {
                        github.issues.createComment({
                            issue_number: context.issue.number,
                            owner: context.repo.owner,
                            repo: context.repo.repo,
                            body: commentText
                        })
                    } catch (error) {
                        console.log('error getting yaml', error)
                    }
                }
                else {
                    try {
                        github.issues.createComment({
                            issue_number: context.issue.number,
                            owner: context.repo.owner,
                            repo: context.repo.repo,
                            body: 'comment from the else'
                            })
                        } catch (error) {
                            console.log("error", error)
                        }
                }
            })
        }
    })
  }