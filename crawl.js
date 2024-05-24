import { JSDOM } from 'jsdom'

function normalizeURL(url)
{
    const tempURL = new URL(url);
    let path = `${tempURL.host}${tempURL.pathname}`
    if (path.slice(-1) === '/')
    {
        path = path.slice(0, -1)
    }
    return path
}

function getURLsFromHTML(htmlBody, baseURL)
{
    const dom = new JSDOM(htmlBody);
    const anchorArray = dom.window.document.querySelectorAll("a");
    const foundArray = []
    for (const anchor of anchorArray)
    {
        if (anchor.hasAttribute('href'))
        {
            let href = anchor.getAttribute('href')
            try
            {
                href = new URL(href, baseURL).href
                foundArray.push(href)
            }
            catch(err)
            {
                console.log(`${err.message}: ${href}`)
            }
        }
    }
    return foundArray
}

async function crawlPage(baseURL, currentURL = baseURL, pages = {})
{
    try 
    {
        const baseURLObj = new URL(baseURL)
        const currentURLObj = new URL(currentURL)
        if (baseURLObj.hostname != currentURLObj.hostname)
        {
            return pages
        }
        const normalized = normalizeURL(currentURL)
        if (pages[normalized] > 0)
        {
            pages[normalized]++
            return pages
        }
        pages[normalized] = 1
        const currentHTML = await fetchHTML(currentURL)
        const nextURLs = getURLsFromHTML(currentHTML, baseURL)
        for (const nextURL of nextURLs)
        {
            pages = await crawlPage(baseURL, nextURL, pages)
        }
        return pages
    }
    catch (error)
    {
        console.log(`Could not crawl page: ${error.message}`)
        return pages
    }
}

async function fetchHTML(url)
{
    try
    {
        const response = await fetch(url)
        if (response.status >= 400)
        {
            console.log(`Error encountered when crawling URL: HTTP ${response.status}`)
            return
        }
        const contentType = response.headers.get('content-type')
        if (!contentType || !contentType.includes('text/html'))
        {
            console.log(`Invalid content type encountered: ${contentType}`)
            return
        }
        return response.text()
    } 
    catch (error) 
    {
        console.log(`Could not crawl URL ${url}: ${error.message}`)
        return
    }
}
export { normalizeURL, getURLsFromHTML, crawlPage }