const CONFIG_FILE = './variables-to-check.yaml'
const yaml = require('js-yaml')
const fs = require('fs')

// const getExistingComments = async (github, owner, repo, pullNumber) => {
//     const allComments = await github.pulls.listComments({
//       owner,
//       repo,
//       number: pullNumber
//     })
//     return allComments.data.filter((comment) => {
//         return comment.user.login === BOT_NAME
//       })
// }

// const commentAlreadyExists = (comments, position, potentialCommentText) => {
//     return !!comments.find((comment) => {
//         return comment.position === position && comment.body === potentialCommentText;
//     })
// }
"look at me! I changed!"

module.exports = async ({github, context}) => {
    const ownerRef = context.repo.owner;
    const repoRef = context.repo.repo;
    const commitIdRef = context.sha
    const pullNumberRef = context.payload.number
// 
tacoCat
    console.log("CONTEXT STUFF", context)

   function getValues() {
        try {
            let config = yaml.load(fs.readFileSync(CONFIG_FILE, 'utf-8'))
            // console.log('Docssss ', config)
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
    // console.log("DIFF HERE YO ", file)


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

                    // console.log("Commenting " + commentText + addedLine + "position : ", position)
                    try {
                        github.pulls.createReviewComment({
                            owner: ownerRef,
                            repo: repoRef,
                            body: (commentText + addedLine),
                            pull_number: pullNumberRef,
                            commit_id: commitIdRef,
                            path: "README.md",
                            line: 3
                        })
                    } catch (error) {
                        console.log('error getting yaml', error)
                    }
                }
            })
        }
      position += 1;
    }
    )
  }

//   await github.rest.pulls.createReviewComment({
//     body: '👋 👋 👋 test',
//     owner: context.repo.owner, 
//     repo: context.repo.repo,
//     pull_number: context.payload.number, 
//     commit_id: context.payload.sha,
//     path: "README.md",
//     line 3