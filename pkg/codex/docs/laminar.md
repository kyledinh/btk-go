# Laminar for ScalaJS

## Usage

build.sbt
```
"com.raquo" %%% "laminar" % "0.13.1"  // Requires Scala.js >= 1.5.0
"com.raquo" %%% "airstream" % "<version>"
```

import
```
import com.raquo.api.L.{*, given}  // Scala 3
```

<br><hr><br>

## Laminar bootstrap example

```scala
import com.raquo.laminar.api.L.{*, given} 
import org.scalajs.dom

@main
def MyApp(): Unit = {
    renderOnDomContentLoader(
        dom.document.querySelector("#app"), Main.appElement())
}

object Main {
    def appElement(): Element = {
        div(
            h1("Hello Scala JS and Laminar"),
            a(href := "https://someurl.com",
                target := "_blank", "Documentation Link"),
        )
    }
}
```

## Render Table Row example 

```scala
  def renderDataTable(): HtmlElement = {
    table(
      thead(
        tr(th("Label"), th("Value"), th("Action")),
      ),
      tbody(
        children <-- dataSignal.split(_.id) { (id, initial, itemSignal) =>
          renderDataItem(id, itemSignal)
        }
      ),
      tfoot(
        tr(td(button("➕", onClick --> (_ => dataVar.update(data => data :+ DataItem()))))),
      ),
    )
  }

  def renderDataItem(id: DataItemID, item: Signal[DataItem]): HtmlElement = {
    val labelUpdater = dataVar.updater[String] { (data, newLabel) =>
      data.map(item => if item.id == id then item.copy(label = newLabel) else item)
    }

    val valueUpdater = dataVar.updater[Double] { (data, newValue) =>
      data.map(item => if item.id == id then item.copy(value = newValue) else item)
    }

    tr(
      td(inputForString(item.map(_.label), labelUpdater)),
      td(inputForDouble(item.map(_.value), valueUpdater)),
      td(button("🗑️", onClick --> (_ => dataVar.update(data => data.filter(_.id != id))))),
    )
  }
```


<br><hr><br>

## Observables

## Vars and Signals
## Modifiers  

<br><hr><br>
## References
- [https://github.com/raquo/Laminar/](https://github.com/raquo/Laminar/) - Under [MIT License](https://github.com/raquo/laminar/blob/master/LICENSE.md).
- [Laminar Official Docs](https://laminar.dev/documentation)
- [Laminar - Smooth UI Youtube](https://www.youtube.com/watch?v=L_AHCkl6L-Q)
- [Slides for Laminar Video](https://docs.google.com/document/d/12dNCHj6QwWCO4BPuFTppn2UcP1rB1gYHly-RGHwoXEI/edit)
- [Xebia Lamiar YouTube Video](https://www.youtube.com/watch?v=UePrOa_1Am8)
- [Git Hub Repo for Xebia Lamiar YouTube Video](https://github.com/sjrd/scalajs-sbt-vite-laminar-chartjs-example)