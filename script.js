const CONFIG_FILE = './variables-to-check.yaml'
const yaml = require('js-yaml')
const fs = require('fs')

const getExistingComments = async (github, owner, repo, pullNumber) => {
    const allComments = await github.pulls.listComments({
      owner,
      repo,
      number: pullNumber
    })
    return allComments.data.filter((comment) => {
        return comment.user.login === BOT_NAME
      })
}

const commentAlreadyExists = (comments, position, potentialCommentText) => {
    return !!comments.find((comment) => {
        return comment.position === position && comment.body === potentialCommentText;
    })
}


module.exports = async ({github, context}) => {
   console.log("MADE IT TO SCRIPTS print", context)
   
   function getValues() {
        try {
            let config = yaml.load(fs.readFileSync(CONFIG_FILE, 'utf-8'))
            console.log('Docssss ', config)
            return config
        } catch (error) {
            console.log('error getting yaml', error)
    }}
    let config = getValues()

    const matchesPattern = (pattern, text) => {
        const regex = new RegExp(pattern);
        return !!text.match(regex)
      };

    const diff_url = context.payload.pull_request.diff_url
    const files = await github.request(diff_url)
    
    const file = files.data
    console.log("DIFF HERE YO ", file)


    // const alreadyHereComments = await getExistingComments(
    //     context.github,
    //     context.repo.owner, 
    //     context.repo.repo,

    // )


    let position = 0
    file.split("\n").forEach((line) => {
        // console.log('made it inloop', line)
        if (line.startsWith("+")){
            const addedLine = line.slice(1);
            const matches = config.matches;

            matches.forEach((match) => {
                const commentText = match.comment ? match.comment : config.defaults.comment;

                if (matchesPattern(match.regex, addedLine)){
                    console.log("Commenting " + commentText + addedLine + "position : ", position)
                    try {
                        github.issues.createComment({
                            issue_number: context.issue.number,
                            owner: context.repo.owner,
                            repo: context.repo.repo,
                            body: (commentText + addedLine)
                        })
                    } catch (error) {
                        console.log('error getting yaml', error)
                    }
                }
                else {
                    console.log('no match here: ', addedLine)
                }
            })
        }
      position += 1;
    })
  }