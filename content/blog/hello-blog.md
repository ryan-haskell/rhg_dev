---
title: "Hello, blog!"
subtitle: "the first post"
description: "New year, new blog! Read about why I chose Hugo for my blog redesign, and what my goals are for the new site."
image: "art"
date: 2022-01-09T01:24:24-06:00
---
<!-- 
- I redesign my blog a lot
- Have a hard time writing actual posts
- i want to change that
- previously rolled my own blog platform
- gonna try to use hugo instead
- focus on how easy it is to add content
- try to become a better writer -->

### Oops, I did it again...

I redesigned my blog site. To be honest, the dopamine rush of the new pretty gradients __almost__ distracted me from the fact that I have made more blogs than blog posts in my lifetime.

Clearly, I am more confident in my ability to create a website than in writing content for the internet. But why is that? I wasn't born with the power to implement blogs. That confidence came from __lots of practice__ coding things.

If I want to become a better writer, I'll need to build up that confidence through actually creating written content.

### Breaking the redesign cycle

With previous website implementations, I made my job as a content author more difficult than necessary. My undying love for [Elm](https://elmlang.org) led me to roll my own custom blogging platform. 

Because of the quality of this custom tool, there were too many manual steps I had to take to publish a new blog post:

1. Create a new markdown file
1. Update the blog list page to include it
1. Make sure the sitemap includes the new URL

This reminds me of a specific experience I feel when adding a feature to an old codebase after a few months. When I build a game or app with Elm, I know I will be able to hop back in and be able to complete my task. On the other hand, when I come back to one of my old JS projects, I spend a lot of time wondering how the heck I wired everything together.

Was this using Gulp or Parcel? Did I use SCSS or CSS-in-JS? Did I use TypeScript? What kind of patterns was I using at this time? All these questions __lower my drive__ to hop in and get things done.

Choosing to roll my own half-baked blogging platform made me feel those insecurities when it came time to write a post. It's easier to write an article on _Medium_, _dev.to_, or even better: to not write anything at all.

If I want to get more practice, I should make the process of writing as impediment-free as I can.

### Choosing Hugo

The website you are on is built with [Hugo](https://hugo.io). Here are a few o fthe characteristics that made Hugo the best choice for my writing goals:

1. __Markdown to HTML__ - I write my blog posts in `.md` files and hugo maps them to HTML and CSS at build-time. This allows me to author content in VS code, GitHub, or a CMS like [Forestry](https://forestry.io).
1. __Free and easy hosting__ - Because the website is static, I can host it for free on [Netlify](https://netlify.com). It connects to my GitHub repo, and deploys when I update markdown or code!
1. __Server-rendered & SEO ready__ - I have full control over how content appears when I share it via Twitter, Discord, Slack, etc. Hugo has the ability to render `<meta>` tags and generates a sitemap for me.
1. __Pre-built archetypes__ - Hugo allows me to create templates called "archetypes" and scaffold new posts with the `hugo new` command. This makes it easy to track what information I need to fill in for each new post on the site.
1. __Blazing fast site builds.__ – As I add more pages, I won't need to worry about slow build times or laggy feedback loops
1. __More than one guy uses it__ - Even though I'm pretty tight with the guy who made my previous blog platform (again, literally me), that dude has a questionable memory. It's nice to have a tool with extensive docs and a community of users.


### Wish me luck!

I'm hoping the ease of content entry (and how much control I have over its presentation) encourages me to share more content online.

My passion for building apps, designing games, and sharing them with a URL is why I got excited about web development in the first place! I'm grateful for my ability to create these experiences, and I hope to get better at sharing them as I go.

__Thank you__ for taking the time to read my first post! There will be many more to come, but if you are reading this in January 2022– you should mess around with some of the games in [my arcade](/arcade)!

