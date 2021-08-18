module.exports = async ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")
    // console.log('Context print ', context)

    const matchesPattern = (pattern, text) => {
        const regex = new RegExp(pattern);
        return !!text.match(regex)
      };

    const diff_url = context.payload.pull_request.diff_url
    const files = await github.request(diff_url)
    console.log("DIFF2: ", files)

    files.data.forEach((file) => {
        console.log("FILE ", file) 
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