function printReport(pages)
{
    console.log("The report is starting")
    console.log("======================")
    const sortedPages = sortPages(pages)
    for (const sortedPage of sortedPages)
    {
        const url = sortedPage[0]
        const count = sortedPage[1]
        console.log(`Found ${count} internal links to ${url}`)
    }
}

function sortPages(pages)
{
    const pagesArray = Object.entries(pages)
    pagesArray.sort((pageA, pageB) => {
        if (pageA[1] === pageB[1])
        {
            return pageA[0].localeCompare(pageB[0])
        }
        return pageB[1] - pageA[1]
    })
    return pagesArray
}

export { printReport }