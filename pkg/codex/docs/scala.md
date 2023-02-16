# Scala Docs 

## Generate a new Scala 3 Project 
- `sbt new scala/scala3.g8` 

Generates:

```
$ tree scala-3-project-template 
scala-3-project-template
├── build.sbt
├── project
│   └── build.properties
├── README.md
└── src
    ├── main
    │   └── scala
    │       └── Main.scala
    └── test
        └── scala
            └── Test1.scala

```

## Sample `sbt.build` for testing 
```
name := "HelloScalaTest"
version := "0.1"
scalaVersion := "3.2.0"

libraryDependencies ++= Seq(
  "org.scalatest" %% "scalatest" % "3.2.9" % Test
)
```

## References
- [API Docs](https://www.scala-lang.org/api/3.2.2/)
- [Scala 3 Book](https://docs.scala-lang.org/scala3/book/introduction.html)
- [Playground](https://scastie.scala-lang.org/)
- [ZIO 2.x](https://zio.dev/overview/getting-started)