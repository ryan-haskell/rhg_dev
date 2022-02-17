---
title: "GraphQL and Elm"
description: "Designing a frontend Elm client that evolves with your GraphQL API."
tags: [elm]

subtitle: "Making inputs that don't bust your face."
image: "hug"

date: 2022-02-17T16:16:33-06:00
---


### Overview

With GraphQL, your queries can accept **required** or **optional** arguments. If you are designing a GraphQL library for Elm, an intuitive choice would be to use “records” for required arguments, and a “list” for optional arguments.

This allows the compiler to let the developer know about missing required fields in the record, and easily omit optional arguments with `[]` if they choose.

Let’s look at an example of what that API would look like in practice:

```elm
{-| 
    Login's input has a required `email` and `password`,
    but no optional fields.
-}
loginInput : GQL.Query.Login.Input
loginInput =
    GQL.Query.Login.Input.new
        { email = "example@email.com"
        , password = "secret1234"
        }

{-| 
    EditUser's input has an optional `email`, `password`, and `bio`,
    but no required fields.
-}
editUserInput : GQL.Query.EditUser.Input
editUserInput =
    GQL.Query.EditUser.Input.new
        [ GQL.Query.EditUser.Input.email "example.two@email.com"
        , GQL.Query.EditUser.Input.password "otherSecret1234"
        , GQL.Query.EditUser.Input.null.bio
        ]

{-| 
    SignUp's input has a required `email` and `password` field,
    and an optional `fullName` field.
-}
signUpInput : GQL.Query.SignUp.Input
signUpInput =
    GQL.Query.SignUp.Input.new
        { email = "example@email.com"
        , password = "secret1234"
        }
        [ GQL.Query.SignUp.Input.fullName "Johnny Smith"
        ]
```

### Problem #1: Inconsistent types for Input functions

Depending on what kind of arguments the query takes, the *type signature* for each input function will vary. As someone learning the library, this might confuse or intimidate you:

```elm
-- 1. Takes a record
GQL.Query.Login.Input.new :
  { email : String, password : String }
  -> Input

-- 2. Takes a list
GQL.Query.EditUser.Input.new :
  List Optional
  -> Input

-- 2. Takes a record and a list
GQL.Query.SignUp.Input.new :
  { email : String, password : String }
  -> List Optional
  -> Input
```

#### A more consistent alternative

We can make these argument consistent by **always** requiring a record and a list input. That would lead to developers needing to type things like this:

```elm
-- Always include an empty [], even if there are no optional fields
GQL.Query.Login.Input.new
    { email = "example@email.com"
    , password = "secret1234"
    }
    []

-- Always include an empty {}, even if there are no required fields
GQL.Query.EditUser.Input.new
    {}
    [ GQL.Query.EditUser.Input.email "example.two@email.com"
    , GQL.Query.EditUser.Input.password "otherSecret1234"
    , GQL.Query.EditUser.Input.null.bio
    ]
```

The rest of this post will assume we do not decide to make developers use `{}` or `[]` where there are no required / optional fields to provide. 

Instead, we will refer to the first API design shown. This is the one popularized by the [@dillonkearns/elm-graphql](https://github.com/dillonkearns/elm-graphql) package.

### Problem #2: APIs change over time!

As the needs for your app evolve over time, so will your API! There are 2 kinds of changes that you might make to your GraphQL API:

1. **Breaking changes** 💔
*( Existing frontend code will need to be updated )*
    - Adding a new required input
    - Making an optional input required
2. **Compatible changes** ✅
*( Existing frontend code will continue to work )*
    - Adding a new optional input
    - Making a required input optional
    - Removing a required input
    - Removing an optional input

Although GraphQL encourages only creating compatible changes for your API, you may still encounter API changes that break your existing frontend Elm application.

However, because your Elm frontend is so resilient, the compiler will walk you through all of those breaking API changes! You’ll need to do a bit of work, but when it compiles, you’ll match the server’s specified GraphQL schema!

So what’s the issue? If we use the “Record/List” pattern above, our Elm code will break- even for **compatible** changes! Here’s how that can happen:

#### 1. Adding an optional field for the first time breaks things!

Adding an optional field should be a compatible change, right? Well *sometimes* it can break your Elm code. 

Let’s say our Login form now takes an optional `phoneNumber` field:

```elm
-- BEFORE
new : { email : String, password : String } -> Input

-- AFTER adding an optional field
new : { email : String, password : String } -> List Optional -> Input
```

```elm
GQL.Query.Login.Input.new
    { email = "example@email.com"
    , password = "secret1234"
    }
    -- Need to add [] here for the compiler to be happy!
```

Even though the new field is optional, the Elm compiler will require you to update all existing instances of `GQL.Query.Login.Input.new` because of the new function argument.

#### 2. Removing a required input breaks things!

Because of our choice to use an Elm record for required fields, anytime we remove a field, we change the shape of that record.

For example, if we removed the `password` field for our `Login` input, this would break all instances of using the `Login` input!

```elm
-- BEFORE
GQL.Query.Login.Input.new :
  { email : String, password : String }
  -> Input

-- AFTER
GQL.Query.Login.Input.new :
  { email : String }
  -> Input
```

```elm
GQL.Query.Login.Input.new
    { email = "example@email.com"
    , password = "secret1234" -- Password needs to deleted!
    }
```

( This problem is actually kind of nice, because it makes us clean up dead code in our application! )

#### 3. Removing the last optional input breaks things!

Earlier, we added an optional `phoneNumber` field to the `Login` input. If we decided to remove that field later on, it could change the type signature of our function. This is because `List Optional` is no longer an argument:

```elm
-- BEFORE
new : { email : String, password : String } -> List Optional -> Input

-- AFTER removing the optional field
new : { email : String, password : String } -> Input
```

```elm
GQL.Query.Login.Input.new
    { email = "example@email.com"
    , password = "secret1234"
    }
    [] -- Need to delete [] here for the compiler to be happy!
```

#### 4. Making a required input optional breaks things!

If we decide password is optional instead of required, we have to change how we pass the `password` value through, because it no longer exists on our record.

Instead we have to use the `GQL.Query.Login.Input.password` function to add it in an optional way:

```elm
-- BEFORE
GQL.Query.Login.Input.new :
  { email : String, password : String }
  -> Input

-- AFTER
GQL.Query.Login.Input.new :
  { email : String }
  -> List Optional
  -> Input
```

```elm
GQL.Query.Login.Input.new
    { email = "example@email.com"
    , password = "secret1234" -- Password needs to deleted!
    }
    [] -- ...and moved into this list!
```

### Designing a better input

My goal is to eliminate the two problems identified above, by using a different strategy for creating GraphQL inputs with Elm. This design will reduce the amount of times our code breaks for “compatible changes” to the API- and provide a consistent experience, regardless of the amount of required vs. optional inputs you have to pass in.

Let’s take a look at the 3 examples from before using the new GraphQL Input pattern:

```elm
{-| Login's input has a required email and password, but no optional fields. -}
loginInput : GQL.Query.Login.Input
loginInput =
    GQL.Query.Login.Input.new
        |> GQL.Query.Login.Input.email "example@email.com"
        |> GQL.Query.Login.Input.password "secret1234"

{-| EditUser's input has an optional email and password, but no required fields -}
editUserInput : GQL.Query.EditUser.Input
editUserInput =
    GQL.Query.EditUser.Input.new
        |> GQL.Query.EditUser.Input.email "example@email.com"
        |> GQL.Query.EditUser.Input.password "secret1234"
        |> GQL.Query.EditUser.Input.null.bio

{-| SignUp's input has a mix of required and optional fields. -}
signUpInput : GQL.Query.SignUp.Input
signUpInput =
    GQL.Query.SignUp.Input.new
        |> GQL.Query.SignUp.Input.email "example@email.com"
        |> GQL.Query.SignUp.Input.password "secret1234"
        |> GQL.Query.SignUp.Input.fullName "Johnny Smith"
```

With this design, required fields are provided the **same way** we provide optional fields. This means that when we make a field go from “required” to “optional”, it’s a compatible change in our Elm application!

### How does it work?

There’s a cool presentation out there by [Jeroen Engels](https://github.com/jfmengels) about an Elm technique called the “**Phantom Builder Pattern**” in Elm. We can use this pattern to enforce required fields without changing how we pass them into our GraphQL inputs.

Here’s a link to [that YouTube video](https://www.youtube.com/watch?v=Trp3tmpMb-o)!

Let’s take a look at an example the generated `GQL.Query.SignUp.Input` module, because it contains a mix of required and optional fields:

```elm
module GQL.Query.SignUp.Input exposing
    ( Input, new
    , email, password
    , fullName
    , null
    )

import GQL.Internals
import GQL.Internals.Input

-- The type variable `missing` is only used in the annotation.
type Input missing
    = Input (List GQL.Internals.Input)

-- The `new` function creates an empty input, and sets the 
-- type variable to an "extensible record", describing which
-- fields are required.
--
-- ( Note that the optional field `fullName` is not included! )
new : Input { missing | email : String, password : String }
new =
    Input []

-- Because `email` is a required field, it changes the signature of
-- `Input { missing | email : String }` into `Input missing`
email : String -> Input { missing | email : String } -> Input missing
email value (Input args) =
    Input (GQL.Internals.Input.string "email" value :: args)

-- Because `password` is also required, it works in a similar way to
-- the `email` function above.
password : String -> Input { missing | password : String } -> Input missing
password value (Input args) =
    Input (GQL.Internals.Input.string "password" value :: args)

-- Because `fullName` is optional, it has no effect on the type signature.
-- The `Input missing` is still `Input missing`
fullName : String -> Input missing -> Input missing
fullName value (Input args) =
    Input (GQL.Internals.Input.string "fullName" value :: args)

-- All optional fields also receive an entry in this `null` record.
-- This makes it easy to explicitly declare an input as "null" for
-- your GraphQL input.
null :
    { fullname : Input missing -> Input missing
    }
null =
    { fullname = \(Input args) -> Input (GQL.Internals.Input.null "fullName" :: args)
    }
```

What do all these weird type signatures mean in practice? Really helpful error messages! If you forget to include a required field, you will see something like this:

```elm
-- TYPE MISMATCH ------------------------------------- ./src/Main.elm

Something is off with the body of the `input` definition:

28|     input : GQL.Query.SignUp.Input
29|     input =
30|         GQL.Query.SignUp.Input.new
31|             |> GQL.Query.SignUp.Input.email "example@email.com"
                ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This value is a:

    GQL.Query.SignUp.Input.Input { missing | password : String }

But the type annotation on `input` says it should be:

    GQL.Query.SignUp.Input
```

The Elm compiler will underline and highlight the missing required fields. By choosing to use actual types in the “extensible record” type variable, we also provide helpful developers information about what their GraphQL API is expecting. 

In this example, we see that `password` field is a `String`.

### But how “compatible” is it?

Before we saw a list of breaking and compatible changes. Let’s compare this new design with the “Record/List API” from before, as well as an ideal (but undiscovered) “Ideal API”.

__Emoji Legend:__

- ✅  = Never leads to compiler errors
- 💔  = Always leads to compiler errors
- 🤷‍♂️  = Sometimes leads to compiler errors

|  | Record / List API | Input Builder API | Ideal API |
| :-- | --- | --- | --- |
| __Breaking Changes__ |  |  |  |
| Adding a new required input | 💔 | 💔 | 💔 |
| Making an optional input required | 💔 | 🤷‍♂️ | 🤷‍♂️ |
| __Compatible Changes__ |  |  |  |
| Adding a new optional input | 🤷‍♂️ | ✅ | ✅ |
| Making a required input optional | 💔 | ✅ | ✅ |
| Removing a required input | 💔 | 🤷‍♂️ | ✅ |
| Removing an optional input | 🤷‍♂️ | 🤷‍♂️ | ✅ |

Take a look at the characteristics of the Input Builder API.

For one of the breaking changes, your frontend code might still work. If you are *already* providing an optional field, and it becomes required– there won't be any compiler errors to fix.

For two compatible changes, the new “Input Builder API” can potentially break. The two 🤷‍♂️ icons you see in the table above will only occur if your code was referencing a removed field.

In these scenarios, the fix is to delete the lines. It will **never** involve moving things around or parsing strange type errors from an added/removed function input.

### That's it!

I hope this article has given you a glimpse into the Elm API design process. The solution isn't perfect, but hopefully it shows how we can still get nice error messages, consistency, and compatibility– even with a strongly typed language like Elm!