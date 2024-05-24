import { crawlPage } from './crawl.js'
import { printReport } from './report.js'

async function main() 
{
    if (process.argv.length === 3)
    {
        const baseURL = process.argv[2]
        console.log(`Starting the crawler at url: ${process.argv[2]}`);
        const pages = await crawlPage(baseURL);
        printReport(pages);
    }
    else
    {
        console.log("Invalid command line arguments - only provide the baseURL");
        return
    }
}

main()