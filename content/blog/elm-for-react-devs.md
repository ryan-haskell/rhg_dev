---
title: "Elm for React Devs"
description: "Are you a React developer? Learn Elm with examples designed for you!"
tags: [elm,react]

subtitle: "They aren't so different!"
image: "react"

date: 2023-03-22T11:29:27-05:00
draft: false
---


When I was first learning [Elm](https://elm-lang.org), it was helpful for me to see side-by-side examples alongside ones from the JavaScript frameworks I was already familiar with.

Recently, the React community released a beautiful guide at [React.dev](https://react.dev/learn) to help beginners learn the framework. This post uses the code examples from their new "Quick Start" guide, to help you leverage your React experience to learn more about Elm!

🚨 __Note:__ The goal of this post is __not__ to compare React to Elm, in an attempt to make one look "better" than the other. Both options are great, so build stuff with whichever you prefer!


### Table of contents

- [Components](#components)
- [Displaying data](#displaying-data)
- [Conditional rendering](#conditional-rendering)
- [Rendering lists](#rendering-lists)
- [Responding to events](#responding-to-events)
- [Sharing data between components](#sharing-data-between-components)

### Components

![Demo of basic react app](./1.png)

When writing React components, we use JSX– a hybrid of HTML and JavaScript syntax:

```jsx
function MyButton() {
  return (
    <button>
      I'm a button
    </button>
  );
}

function MyApp() {
  return (
    <div>
      <h1>Welcome to my app</h1>
      <MyButton />
    </div>
  );
}
```

Elm's syntax doesn't looks like HTML, but the concepts are mostly the same. We can define functions like `myButton` or `myApp`, and reuse them alongside the normal HTML tags imported on line 1:

```elm
import Html exposing (..)

myButton =
    button [] [ text "I'm a button" ]

myApp =
    div []
        [ h1 [] [ text "Welcome to my app" ]
        , myButton
        ]
```

🔗 __Demo:__ [Components](https://ellie-app.com/mkdqdz8PYWga1)

### Displaying data

![Demo of the Hedy Lamarr app](2.png)

Displaying data in JSX is easy– we can use the `{}` characters to "escape" from our HTML and provide values from JavaScript:

```jsx
const user = {
  name: 'Hedy Lamarr',
  imageUrl: 'https://i.imgur.com/yXOvdOSs.jpg',
  imageSize: 90,
};

export default function Profile() {
  return (
    <>
      <h1>{user.name}</h1>
      <img
        className="avatar"
        src={user.imageUrl}
        alt={'Photo of ' + user.name}
        style={{
          width: user.imageSize,
          height: user.imageSize
        }}
      />
    </>
  );
}
```

Because Elm doesn't use HTML syntax, you won't need any escape characters to display data. You can reference those values as if you were writing normal JavaScript code:

```elm
import Html exposing (..)
import Html.Attributes exposing (..)

user =
    { name = "Hedy Lamarr"
    , imageUrl = "https://i.imgur.com/yXOvdOSs.jpg"
    , imageSize = 90
    }

profile =
    [ h1 [] [ text user.name ]
    , img
        [ class "avatar"
        , src user.imageUrl
        , alt ("Photo of " ++ user.name)
        , style "width" (String.fromInt user.imageSize)
        , style "height" (String.fromInt user.imageSize)
        ]
        []
    ]
```

🔗 __Demo:__ [Displaying data](https://ellie-app.com/mkdvmmbQW5wa1)

### Conditional rendering

![Conditional rendering demo](3.gif)

In React, you can use JavaScripts `if/else` syntax or the ternary operator to conditionally render your content:

```jsx
let content;
if (isLoggedIn) {
  content = <AdminPanel />;
} else {
  content = <LoginForm />;
}

// USING TERNARY INSTEAD:
// let content = (isLoggedIn) ? <AdminPanel /> : <LoginForm />
```

In Elm, the `if/else` expression always returns its value. That means content will be set, depending on `isLoggedIn`.

For me, it was helpful to think of Elm's `if/else` like it was ternary, but with the readability of JavaScript's `if` statement:

```elm
content =
    if isLoggedIn then
        adminPanel

    else
        loginForm
```

🔗 __Demo:__ [Conditional rendering](https://ellie-app.com/mkdzqxbLX9wa1)


### Rendering lists

![Rendering lists demo](4.png)

All arrays in JavaScript support the `.map` function, which allows React developers to return lists of components like these list items:

```jsx
const products = [
  { title: 'Cabbage', isFruit: false, id: 1 },
  { title: 'Garlic', isFruit: false, id: 2 },
  { title: 'Apple', isFruit: true, id: 3 },
];

export default function ShoppingList() {
  const listItems = products.map(product =>
    <li
      key={product.id}
      style={{
        color: product.isFruit ? 'magenta' : 'darkgreen'
      }}
    >
      {product.title}
    </li>
  );

  return (
    <ul>{listItems}</ul>
  );
}
```

In Elm, values don't have methods, so we use `List.map` with the function and array we are looping over:

```elm
products =
    [ { title = "Cabbage", isFruit = False, id = 1 }
    , { title = "Garlic", isFruit = False, id = 2 }
    , { title = "Apple", isFruit = True, id = 3 }
    ]

shoppingList =
    let
        listItems =
            List.map 
                (\product ->
                    li
                      [ if product.isFruit then 
                          style "color" "magenta"

                        else
                          style "color" "darkgreen"
                      ]
                      [ text product.title ]
                )
                products
    in
    ul [] listItems
```

🔗 __Demo:__ [Rendering lists](https://ellie-app.com/mkdH5fsmycTa1)

__Note:__ Elm functions can use `let/in` to have "locally-scoped" variables like `listItems`. When we nested `listItems` inside of the `let/in` block, it meant that only the code in `shoppingList` can access to the `listItems` variable. This is unlike `products`, which is available to any function in our Elm file.


### Responding to events

![Responding to events demo](5.gif)

In React, you'll use "hooks" like `useState` to track application state. Here's an example of how to do this with a counter and a button:

```jsx
function MyButton() {
  const [count, setCount] = useState(0);

  function handleClick() {
    setCount(count + 1);
  }

  return (
    <button onClick={handleClick}>
      Clicked {count} times
    </button>
  );
}
```

All Elm applications use the same `init/update/view` pattern to keep track of application state.

In our example:
- `init` sets the initial value of the counter
- `update` defines how our count changes from user events
- `view` defines how to render our HTML, based on the latest count

```elm
import Html exposing (..)
import Html.Events exposing (..)

-- INIT

init = 0

-- UPDATE

type Msg = HandleClick

update msg model =
    case msg of
        HandleClick ->
            model + 1

-- VIEW

view model =
    button 
        [ onClick HandleClick ]
        [ text ("Clicked " ++ String.fromInt model ++ " times")
        ]
```

🔗 __Demo:__ [Responding to events](https://ellie-app.com/mkdKQgMk3Pma1)

### Sharing data between components

![Sharing data between components demo](6.gif)

In React, our state is defined within our component functions. If we want to share it, we pass "props" like `count` or `handleClick` into child components like `<MyButton>`:

```jsx
function MyApp() {
  const [count, setCount] = useState(0);

  function handleClick() {
    setCount(count + 1);
  }

  return (
    <div>
      <h1>Counters that update together</h1>
      <MyButton count={count} onClick={handleClick} />
      <MyButton count={count} onClick={handleClick} />
    </div>
  );
}

function MyButton({ count, onClick }) {
  return (
    <button onClick={onClick}>
      Clicked {count} times
    </button>
  );
}
```

In Elm apps, we always use the `init/update/view` pattern. This means we can pass our `count` and `HandleClick` event into our `myButton` function as props from our top-level `view` function:

```elm
import Html exposing (..)
import Html.Events exposing (..)

-- INIT

init = 0

-- UPDATE

type Msg = HandleClick

update msg model =
    case msg of
        HandleClick ->
            model + 1

-- VIEW

view model =
    div []
      [ h1 [] [ text "Counters that update together" ] 
      , myButton { count = model, onClick = HandleClick }
      , myButton { count = model, onClick = HandleClick }
      ]

myButton props =
    button 
        [ onClick props.onClick ]
        [ text ("Clicked " ++ String.fromInt props.count ++ " times")
        ]
```

🔗 __Demo:__ [Sharing data between components](https://ellie-app.com/mkdMsrCMGhPa1)

#### That's all I've got!

I hope this tiny guide was helpful, and that you learned something new today. If you want to learn more about Elm, you should check out [the official guide](https://guide.elm-lang.org) or say hello in the [Elm Slack #beginners channel](https://elm-lang.org/community/slack)

Have a great day! 👋
