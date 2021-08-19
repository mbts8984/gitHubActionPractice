const CONFIG_FILE = './variables-to-check.yaml'

module.exports = async ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")
    // console.log('Context print ', context)

    const matchesPattern = (pattern, text) => {
        const regex = new RegExp(pattern);
        return !!text.match(regex)
      };

    const diff_url = context.payload.pull_request.diff_url
    const files = await github.request(diff_url)
    
    const file = files.data
    const config = await context.config(CONFIG_FILE)

    console.log("CONFIG HERE: ", config)
    
    let position = 0
    file.split("\n").forEach((line) => {
        // console.log('made it inloop', line)
        if (line.startsWith("+")){
            const addedLine = line.slice(1);
            console.log('++++ ', addedLine)
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