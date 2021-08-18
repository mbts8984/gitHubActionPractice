module.exports = ({github, context}) => {
    console.log("MADE IT TO SCRIPTS print")
    console.log('Context print ', context)

    github.issues.createComment({
        issue_number: context.issue.number,
        owner: context.repo.owner,
        repo: context.repo.repo,
        body: '👋 Thanks for reporting!'
      })
    
  }