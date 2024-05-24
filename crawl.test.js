import { normalizeURL, getURLsFromHTML } from './crawl.js'
import { test, expect } from '@jest/globals'

test('normalizeURL https', () => {
  const input = 'https://blog.boot.dev/path'
  const actual = normalizeURL(input)
  const expected = 'blog.boot.dev/path'
  expect(actual).toEqual(expected)
})

test('normalizeURL http', () => {
  const input = 'http://blog.boot.dev/path'
  const actual = normalizeURL(input)
  const expected = 'blog.boot.dev/path'
  expect(actual).toEqual(expected)
})

test('normalizeURL slash', () => {
  const input = 'http://blog.boot.dev/path/'
  const actual = normalizeURL(input)
  const expected = 'blog.boot.dev/path'
  expect(actual).toEqual(expected)
})

test('getURLsFromHTML no urls', () => {
    const inputBody = '<p>Not a URL</p>'
    const inputURL = 'https://blog.boot.dev'
    const actual = getURLsFromHTML(inputBody, inputURL)
    const expected = 0
    expect(actual.length).toEqual(expected)
})

test('getURLsFromHTML one url', () => {
    const inputBody = '<html><body><a href="https://blog.boot.dev">Learn Backend Development</a></body></html>'
    const inputURL = 'https://blog.boot.dev'
    const actual = getURLsFromHTML(inputBody, inputURL)
    const expected = [ 'https://blog.boot.dev/' ]
    expect(actual).toEqual(expected)
})

test('getURLsFromHTML relative to absolute', () => {
    const inputBody = '<html><body><a href="/path/one">Learn Backend Development</a></body></html>'
    const inputURL = 'https://blog.boot.dev'
    const actual = getURLsFromHTML(inputBody, inputURL)
    const expected = [ 'https://blog.boot.dev/path/one' ]
    expect(actual).toEqual(expected)
})

test('getURLsFromHTML multiple URLs', () => {
    const inputBody = '<html><body><a href="/path/one">Learn Backend Development</a><a href="https://google.com">Google</a></body></html>'
    const inputURL = 'https://blog.boot.dev'
    const actual = getURLsFromHTML(inputBody, inputURL)
    const expected = [ 'https://blog.boot.dev/path/one', 'https://google.com/' ]
    expect(actual).toEqual(expected)
})