# Go代码整洁之道

## 前言

编写本文档背后的目的是为Go社区贡献一个资源（并最终作为一个参考），可以帮助开发人员编写出更整洁(清晰)的代码。不管我们是自己在Coding，或者是在更大的团队中Coding。这都是有益于我们每个人。建立良好的案例(规范)用于编写清晰的代码可以确保每个人都可以使用的良好范例，会有效防止无意义的时间浪费在试图去理解和解析其他人（和我们自己的）代码。

> We don’t read code, we decode it - Peter Seibel
>
> 我们不读代码，而是去解析它。

事实上，正如Peter Seibel所说的那样。我们分析代码，老实说无法以某种方式，形状或形式对其进行编码。本文档将成为我们的指导，以确保我们的编码方法是有效的。我们希望我们的代码具有**可行性**，**可读性**和**可维护性**。

作为开发人员,我们有时更倾向于以一种方便的方式去Coding,而不考虑是否是最佳实践;这使得代码评审和测试更加困难。从某种意义上说,我们正在这样Coding,会使其他人更难解析我们的作品。但是我们希望我们的代码是**可行的、可读的和可维护的**。这需要以正确的方式进行Coding,而不是以最简单的方式去Coding。

本文档将首先简要介绍编写清晰代码背后的基础知识，下面我们主要讨论主要关于Go的具体的重构示例。

### **简单说说`gofmt`**

我想用几句话来表达我对`gofmt`的看法。在涉及到`gofmt`这个工具的时候，我不同意很多地方。我更喜欢下划线命名(snake_case)而不是驼峰命名(camelCase)来命名变量，我很喜欢常量用大写，我对括号位置放置也不太同意。如是说：`gofmt`使我们能够有一个编写Go代码的通用标准。所有Go代码看起来都有些相似，它确保Go代码不会变得过于开放，这也是一件非常好的事情。作为开发人员的我可以理解这对Go程序员都有点被`gofmt`限制，尽管对这些规则不太同意。但在我看来，同质代码笔拥有完全的表达自由更重要。

## 内容目录

- [Clean Code简介](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Introduction-to-Clean-Code)
  - [测试驱动开发](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Test-Driven-Development)
  - [命名规则](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Naming)
  - - 注释
      - [函数命名](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Function-Naming)
      - [变量命名](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Variable-Naming)
  - 函数规则
    - [函数长度](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Function-Length)
    - [函数签名](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Function-Signatures)
  - [可变范围](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Variable-Scope)
  - [变量声明](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Variable-Declaration)
- [干净去](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Clean-Go)
  - 返回值
    - [返回定义的错误](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Returning-Defined-Errors)
    - [返回动态错误](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Returning-Dynamic-Errors)
  - [Go中的指针](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Pointers-in-Go)
  - [闭包是函数指针](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Closures-are-Function-Pointers)
  - [Go中的接口](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Interfaces-in-Go)
  - [空的 `interface{}`](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#The-empty-`interface{}`)
- [摘要](https://github.com/Pungyeon/clean-go-article?utm_campaign=The Go Gazette&utm_medium=email&utm_source=Revue newsletter#Summary)

## 代码优化简介

Clean Code，是提升可读和可维护软件的实用概念。Clean Code建立了对代码库的信任，并有助于将出现粗心错误的可能性降至最低。它还可以帮助开发人员保持敏捷性，但由于引入 Bug 的风险增加,随着代码库的扩展这种敏捷性通常会急剧下降。

### 测试驱动开发

测试驱动开发是在短的开发周期或冲刺 (sprint) 中频繁测试代码的做法。它通过邀请开发人员质疑其代码的功能和用途,最终有助于实现代码的清洁。为了简化测试,建议开发人员编写只执行一件事的简短函数。例如,测试(和理解)一个只有 4 行长的函数比 40 行的函数要容易得多。

测试驱动开发包括以下周期:

1. 编写(或执行)测试
2. 如果测试失败,则使其通过
3. 相应地重构代码
4. 重复

在此过程中,测试和重构是相互紧密联系的。随着重构代码的进行也会使其更容易理解或可维护,你需要彻底地测试做出的更改,以确保没有更改函数的行为。随着代码库的增长,这可能非常有用的。

### 命名规范

#### 注释

首先：我想谈谈注释这个话题。不必要的注释是代码味道的最大吐槽。注释通常被添加到代码库中,因为某些内容非常地不明确,因此有必要对其进行解释,以便读者能够了解发生了什么。但情况并非总是如此,注释往往被误用。

在Go中，依据`gofmt`所有公共变量和函数都应该加注释。我认为这绝对是好的，因为它为我们提供了一致的规则来记录我们的代码。但是,我总是想要区分启用自动生成的文档的注释和所有其他注释。对于文档来说,注释注释应该像文档一样编写,它们应该处于高度抽象级别,并且尽可能少地关注代码的逻辑实现。

我这样说是因为还有其他方法来解释代码,并确保代码的编写是可理解和富有表现力的。如果代码不是这两者,有些人认为可以引入注释来解释复杂的逻辑。大多数人根本不会阅读注释,因为它们对审阅代码的体验是非常具有侵入性。

让我们退一步,看看一些具体的例子。以下是不应该对代码进行注释的方式:

```go
// iterate over the range 0 to 9 
// and invoke the doSomething function
// for each iteration
// 译文：在0到9的范围内迭代，并在每个迭代中并调用 doSomething 函数
for i := 0; i < 10; i++ {
  doSomething(i)
}
```

这个注释，就是我所说的教学式注释。这在教程中很常见，它解释了语言的低级功能（或者更一般的编程）。虽然这些评论对初学者可能有帮助,但这些注释在生产代码中是绝对没用的。希望,在开始开发团队工作时,我们不会与那些不了解循环构造等简单内容的程序员协作。作为程序员,我们不应该通过看注释来了解发生了什么 ，我们知道,我们迭代的范围是 0 到 9,因为我们可以简单地读取代码。正如谚语所说:

> Document why, not how. – Venkat Subramaniam
>
> 记录为什么，而不是怎么办。				

遵循这个逻辑，我们现在可以改变我们的注释，解释为什么我们从0到9的范围迭代：

```go
// instatiate 10 threads to handle upcoming work load
// 控制 10 个线程来处理即将到来的工作负载
for i := 0; i < 10; i++ {
  doSomething(i)
}
```

现在,我们了解了为什么我们做一个循环，并可以通过简单地阅读代码来告诉我们正在做什么...是不是有点

这仍然不是我所认为的整洁代码。注释令人担忧,因为假设代码写得很好(但是它不是)。因此,可能没有必要用散文来表达这样的解释。从技术上讲,我们仍在说我们在做什么,而不是我们为什么这样做。我们可以通过使用更有意义的名称，很容易地直接在我们的代码中表达这个"what"：

```go
for workerID := 0; workerID < 10; workerID++ {
  instantiateThread(workerID)
}
```

只需对变量和函数名称进行一些更改,我们就直接在我们的代码中建立了对操作的解释。这一点对读者来说要清楚得多,因为他们不必看注释,然后将散文注释映射到代码中。相反,他们只需要阅读代码，即可了解代码的编写内容。

当然,这是一个相对微不足道的例子。不幸的是,编写清晰和富有表现力的代码并不总是那么容易;随着代码库本身复杂性的增加,它变得越来越困难。你越是练习用这种心态写注释,避免解释你在做什么,你的代码就会变得越干净。

#### 函数命名

现在,让我们继续讨论函数的命名规则。一般规则非常简单:函数越具体,其名称就越通用。换句话说,我们希望从描述一般功能的非常广泛和短的函数名称(如`Run`或`Parse`开始。让我们想象一下,我们正在创建一个配置解析器。遵循此命名规则,我们的顶层抽象可能如下所示:

```go
func main() {
    configpath := flag.String("config-path", "", "configuration file path")
    flag.Parse()

    config, err := configuration.Parse(*configpath)
    
    ...
}
```

我们重点将介绍Parse函数的命名。尽管此函数名称非常简短且通用,但实际上该函数试图实现的目标还是很清楚的。当我们深入一层时,我们的函数命名将变得更加具体：

```go
func Parse(filepath string) (Config, error) {
    switch fileExtension(filepath) {
    case "json":
        return parseJSON(filepath)
    case "yaml":
        return parseYAML(filepath)
    case "toml":
        return parseTOML(filepath)
    default:
        return Config{}, ErrUnknownFileExtension
    }
}
```

在这里,我们清楚地区分了嵌套函数调用与其父函数调用,而不是很具体的。这允许每个嵌套函数调用本身以及在父函数的上下文中有意义。另一方面,如果我们改为命名`parseJSON`替换函数名json,`json`不止代表自己。该功能将在名称中丢失,我们将不再能够判断此函数是解析、创建还是序列化JSON。

请注意，`fileExtension`实际上更具体一点。但是，这是因为此函数的功能实际上非常具体：

```go
func fileExtension(filepath string) string {
    segemnts := strings.Split(filepath, ".")
    return segments[len(segments)-1]
}
```

我们的函数名称中的这种逻辑发展— 从高级抽象到更低级、更具体的抽象,使代码更易于理解和阅读。考虑替代方案:如果我们的最高抽象级别过于具体,那么我们最终将有一个名称,试图覆盖所有基础的名称,如`DetermineFileExtensionAndParseConfigurationFile`这是可怕的难以阅读;我们试图太具体太快,最终混淆读者,尽管我们的意图是试图明确!

#### 变量命名

更有趣的是,变量的情况正好相反。与函数不同，我们的变量应该随着嵌套范围的加深而从更具体到更不具体地命名。

> You shouldn’t name your variables after their types for the same reason you wouldn’t name your pets 'dog' or 'cat'.   – Dave Cheney
>
> 你不应该以变量的类型来命名变量，就像你不应该为你的宠物命名“dog”或“cat”一样。

为什么当我们深入到函数的作用域时,变量名称会变得不那么具体?简单地说,随着变量的范围变小,读者对该变量代表的内容越来越清晰,因此无需进行特定的命名。在前面的函数`fileExtension`的示例中,我们甚至可以将变量`segments`的名称缩短为 s,如果`s`我们想要的话。变量的上下文非常清晰,因此无需用较长的变量名称进一步解释它。另一个好的例子是在嵌套的 for 循环中

```go
func  PrintBrandsInList（brands [] BeerBrand）{
     for  _，b  ：=  range brands { 
        fmt.Println（b）
    } 
}
```

在上面的示例中,变量`b`的范围非常小,我们无需花费任何额外的脑力来记住它的确切表示内容。然而,由于`brands`的范围稍大,它有助于它更具体。在扩展以下函数中的变量范围时,这种区别变得更加明显:

```go
func BeerBrandListToBeerList(beerBrands []BeerBrand) []Beer {
    var beerList []Beer
    for _, brand := range beerBrands {
        for _, beer := range brand {
            beerList = append(beerList, beer)
        }
    }
    return beerList
} 
```

漂亮！这个函数易于阅读。现在,让我们在命名变量时应用相反(即错误)逻辑:

```go
func BeerBrandListToBeerList(b []BeerBrand) []Beer {
    var bl []Beer
    for _, beerBrand := range b {
        for _, beerBrandBeerName := range beerBrand {
            bl = append(bl, beerBrandBeerName)
        }
    }
    return bl
}
```

即使可以确定此函数正在执行的操作,但是变量名的过于简短使得我们在深入研究时很难遵循逻辑。由于我们将短变量名和长变量名混合在一起，因此很可能会陷入完全混乱。

### 清洁功能

现在,我们已经了解了命名变量和函数的一些最佳实践,以及用注释来解释我们的代码,接下来让我们深入了解如何重构函数以使它们更简洁。

#### 函数长度

> How small should a function be?Smaller than that!  – Robert C. Martin
>
> "函数应该有多小?比这个更小!"



编写干净的代码时,我们的主要目标是使我们的代码易于理解。最有效的方法,就是使我们的函数尽可能短小。请务必了解,这不一定是为了避免代码重复。更重要的原因是提高代码理解能力。

它可以帮助在非常高的层次上查看函数的描述,有助于更好地了解这一点:

```go
fn GetItem:
    - parse json input for order id
    - get user from context
    - check user has appropriate role
    - get order from database
```

通过编写短函数(在 Go 中通常是 5~8 行),我们可以创建几乎像上面描述一样自然读取的代码:

```go
var (
    NullItem = Item{}
    ErrInsufficientPrivliges = errors.New("user does not have sufficient priviliges")
)

func GetItem(ctx context.Context, json []bytes) (Item, error) {
    order, err := NewItemFromJSON(json)
    if err != nil {
        return NullItem, err
    }
    if !GetUserFromContext(ctx).IsAdmin() {
	      return NullItem, ErrInsufficientPrivliges
    }
    return db.GetItem(order.ItemID)
}
```

使用较小的函数还可以消除编写代码的另一个可怕的习惯:缩进地狱。**缩进地狱**通常发生在`if`语句链不小心嵌套在函数中时。这使得我们很难解析代码,并且只要发现就应该删除它。缩进地狱在使用`interface{}`和使用类型强制转换时尤为常见:

```go
func GetItem(extension string) (Item, error) {
    if refIface, ok := db.ReferenceCache.Get(extension); ok {
        if ref, ok := refIface.(string); ok {
            if itemIface, ok := db.ItemCache.Get(ref); ok {
                if item, ok := itemIface.(Item); ok {
                    if item.Active {
                        return Item, nil
                    } else {
                      return EmptyItem, errors.New("no active item found in cache")
                    }
                } else {
                  return EmptyItem, errors.New("could not cast cache interface to Item")
                }
            } else {
              return EmptyItem, errors.New("extension was not found in cache reference")
            }
        } else {
          return EmptyItem, errors.New("could not cast cache reference interface to Item")
        }
    }
    return EmptyItem, errors.New("reference not found in cache")
}
```

// **TODO**

这种代码不仅会给其他程序员带来非常糟糕的体验,他们将不得不努力去理解代码的流。如果我们`if`语句中的逻辑扩展,则要找出哪个语句返回什么变得更加困难。不幸的是,在代码中发现这种实现并不罕见。我甚至碰到了一些例子,`if`一个对应`else`的语句在我的监视器的另一页上。在尝试找出函数的作用时,必须上下滚动页面并不理想。即使,我们不需要滚动在我们的页面上,看到`if else`代码示例中的语句,我们仍然滚动与眼睛和保持我们的大脑状态。大多数程序员可以很容易地包含上述函数的这种状态,或者更糟的例子。然而,我们强迫代码的读者使用不必要的脑力。这可能导致读者疲劳,如果我们在整个代码中重复这个错误。不断需要解析上述代码,会使阅读代码变得越来越困难,我们当然要避免。

那么,我们如何清理这个功能呢?幸运的是,它实际上非常简单。在第一次迭代中,我们将尽力确保尽快返回错误。我们希望"将代码向左推",而不是嵌套`if` `else`语句。这是通过返回我们的功能,尽可能尽快处理。

```go
func GetItem(extension string) (Item, error) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return EmptyItem, errors.New("reference not found in cache")
    }

    if ref, ok := refIface.(string); ok {
        // return cast error on reference 
    }

    if itemIface, ok := db.ItemCache.Get(ref); ok {
        // return no item found in cache by reference
    }

    if item, ok := itemIface.(Item); ok {
        // return cast error on item interface
    }

    if !item.Active {
        // return no item active
    }

    return Item, nil
}
```

一旦我们完成了这个,我们可以拆分我们的函数到较小的函数,如前面所述。一个很好的经验法则是:如果`value, err :=`模式在函数中重复多次,则表明我们可以将代码的逻辑拆分为较小的函数。

```go
func GetItem(extension string) (Item, error) {
    if ref, ok := getReference(extension) {
        return EmptyItem, ErrReferenceNotFound
    }
    return getItemByReference(ref)
}

func getReference(extension string) (string, bool) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return EmptyItem, false
    }
    return refIface.(string)
}

func getItemByReference(reference string) (Item, error) {
    item, ok := getItemFromCache(reference)
    if !item.Active || !ok {
        return EmptyItem, ErrItemNotFound
    }
    return Item, nil
}

func getItemFromCache(reference string) (Item, bool) {
    if itemIface, ok := db.ItemCache.Get(ref); ok {
        return EmptyItem, false
    }
    return itemIface.(Item), true
}
```

> 对于生产代码,应该通过返回错误而不是`bool`值来进一步阐述代码。这使得更容易理解错误的来源。但是,由于这些只是示例函数,`bool`值现在就足够了。稍后将更详细地解释更明确地返回错误的示例。

生成的干净版本我们的函数,已经产生了更多的代码行。但是,代码更易于阅读。它以洋葱风格分层,我们可以忽略我们不感兴趣的代码详细信息,并深入了解我们希望了解背后的功能。当我们深入探讨较低级别的功能时,将非常容易理解,因为在这种情况下,我们只需要了解 3-5 行。此示例说明,我们无法从函数的行计数中评分代码的清洁度。第一个函数迭代要短得多。然而,它被人为地短,很难阅读。在大多数情况下,清理代码将首先根据代码行展开现有的代码库。但是,可读性的好处是首选。如果您对此有疑问,请思考您对于以下函数的感觉,该函数也是如此:

```
func GetItemIfActive(extension string) (Item, error) {
    if refIface,ok := db.ReferenceCache.Get(extension); ok {if ref,ok := refIface.(string); ok { if itemIface,ok := db.ItemCache.Get(ref); ok { if item,ok := itemIface.(Item); ok { if item.Active { return Item,nil }}}}} return EmptyItem, errors.New("reference not found in cache")
}
```

当我们讨论这个话题时。在编写这种代码样式时,还会产生很多其他副作用。很显然,它使我们的代码更容易测试。在 4 行(由理智的人编写)的函数上获取 100% 的代码覆盖率要比 400 行的函数容易得多。这是常识。

#### 函数签名

创建良好的函数命名结构,便于阅读和理解代码的意图。缩短函数,有助于理解函数逻辑的内容。清理函数的最后一部分,将是了解函数输入的上下文。有了这个,又来了另一个容易遵循的规则。函数签名,应仅包含一个或两个输入参数。在某些特殊情况下,三个是可以接受的,但我们应该开始考虑重构。就像我们的函数应该只有 5-8 行长的规则一样,这在一开始似乎相当极端。然而,我觉得这个规则更明显地是正确的。

例如,从 RabamQ 简介教程到其 Go 库,请执行以下功能:

```
q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
```

函数`QueueDeclare`采用六个输入参数,这是相当极端的。由于注释,上述代码有些可以理解,但如前所述:注释应替换为描述性代码。一个很好的原因是,没有任何东西阻止我们调用`QueueDeclare`功能而不进行注释,使其如下所示:

```
q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
```

现在,不看前面的代码,尝试记住第四个和第五`false`表示什么。这是不可能的,我们不可避免地会在某个时候忘记。这可能导致代价高昂的错误和难以纠正的错误。错误甚至可能通过不正确的注释发生。想象一下,标记错误的输入参数。纠正这个错误,将难以纠正,尤其是当对代码的熟悉度随着时间的推移而恶化或开始程度很低时。因此,建议用"选项"`struct`替换这些输入参数:

```go
type QueueOptions struct {
    Name string
    Durable bool
    DeleteOnExit bool
    Exclusive bool
    NoWait bool
    Arguments []interface{} 
}

q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
    Durable: false,
    DeleteOnExit: false,
    Exclusive: false,
    NoWait: false,
    Arguments: nil,
})
```

这解决了省略注释或意外标记变量错误的问题。当然,我们仍然可以将属性与错误的值混淆,但在这些情况下,确定我们的错误在代码中的位置会容易得多。属性的顺序也不再重要,因此输入值的顺序不正确,不再是一个担心。此技术的最后一个好处是,我们可以使用我们的 Option`struct`来推断函数输入参数的默认值。在 Go 中声明结构时,所有属性都初始化为默认值。这意味着,我们的`QueueDeclare`选项实际上可以以下列方式调用:

```go
q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
})
```

其余值通过初始化为默认`false`值(`Arguments`除外),作为接口的默认值为`nil`我们不仅更安全,我们更清楚地知道我们的意图,在这种情况下,我们实际上可以编写更少的代码。这是一场全面的胜利。

最后一点,不是总是可以更改函数签名。与本例中一样,我们无法控制`QueueDeclare`函数签名,因为这是来自 RabamQ 库的。这不是我们的代码,我们不能改变它。但是,我们可以包装这些函数,以适应我们的目的:

```go
type RMQChannel struct {
    channel *amqp.Channel
}

func (rmqch *RMQChannel) QueueDeclare(opts QueueOptions) (Queue, error) {
    return rmqch.channel.QueueDeclare(
        opts.Name,
        opts.Durable,
        opts.DeleteOnExit,
        opts.Exclusive,
        opts.NoWait,
        opts.Arguments, 
    )
} 
```

基本上,我们创建一个新的`RMQChannel``amqp.Channel`类型,具有`QueueDeclare`方法。然后,我们创建自己的此方法版本,它基本上只调用旧版本的 RabamQ 库函数。我们的新方法具有之前描述的所有优点,我们实现了这一点,实际上无需更改 RabamQ 库中的任何代码。

稍后讨论`interface{}`时,我们将使用包装函数的概念来引入更干净、更安全的代码。.

兔子MQ围棋教程

### 可变范围

现在,让我们回到一步,回到编写较小函数的想法。这有另一个不错的副作用,我们在上一章中没有介绍:编写较小的函数通常可以使用更持久的可变变量消除。使用全局变量编写代码,是过去的做法,它不属于干净的代码。为什么会这样?那么,使用全局变量的问题是,我们使得程序员很难理解变量的当前状态。如果变量是全局的且可变的,那么,根据定义,它的值可以由代码库的任何部分更改。您不能保证此变量将是一个特定的值...这对每个人都是一个头痛的问题。这是一个微不足道的问题,当代码库扩展时,这个问题会加剧。让我们看一个简短的示例,了解更大的范围(不是全局)变量如何导致问题。

范围较大的变量,还引入了变量阴影问题,如来自名为:Golang 范围问题的文章的代码所示 -[`Golang scope issue - A feature bug: Shadow Variables`](https://www.translatoruser-int.com/translate?&to=zh-CHS&csId=24f5af06-cfe4-44a8-9412-da1a7e7ea8c9&usId=7f43ce80-5598-418d-ac38-b19975567286&dl=en&ac=true&dt=2019%2f6%2f23 10%3a19&h=GjQkqwMvbs82O2M-CFSMrZ1cMnPp8Gn3&a=https%3a%2f%2fidiallo.com%2fblog%2fgolang-scopes):

```go
func doComplex() (string, error) {
    return "Success", nil
}

func main() {
    var val string
    num := 32

    switch num {
    case 16:
    // do nothing
    case 32:
        val, err := doComplex()
        if err != nil {
            panic(err)
        }
        if val == "" {
            // do something else
        }
    case 64:
        // do nothing
    }
    
    fmt.Println(val)
}
```

此代码的问题,从快速浏览,似乎`var val string`值,应打印出来为:`Success`到`main`函数的末尾。不幸的是,情况并非如此。其原因是,行:

> 瓦尔, 错误 := 多复杂()

这将在 switch 大小`val`写`32`作用域中声明一个新的变量 val,并且与`main`的第一行中声明的变量无关。当然,可以说,Go语法有点棘手,我不一定不同意,但手头有一个更棘手的问题。`var val string`声明为可变的、范围很广的变量,是完全不必要的。如果我们执行**非常简单**的重构,我们将不再有此问题:

```go
func getStringResult(num int) (string, error) {
    switch num {
    case 16:
    // do nothing
    case 32:
       return doComplex()
    case 64:
        // do nothing
    }
    return "" 
}

func main() {
    val, err := getStringResult(32)
    if err != nil {
        panic(err)
    }
    if val == "" {
        // do something else
    }
    fmt.Println(val)
}
```

重构后,`val`不再发生突变,范围已缩小。同样,请记住这些函数非常简单。一旦这种代码样式成为更大更复杂的系统的一部分,就不可能找出错误发生的原因。我们不希望这种情况发生。不仅因为我们通常不喜欢软件中发生的错误,而且我们的同事和我们自己也不尊重我们,因为我们可能浪费了彼此的活,不得不调试这种类型的代码。让我们自己承担责任,而不是责怪 Go 中的可变声明语法。

在一侧,如果`// do something else`是另一个尝试突变`val`变量。我们应该提取其中的任何逻辑作为函数,以及它的前一部分。这样,我们可以返回一个新值,而不是延长变量的突变范围:

```go
func getVal(num int) (string, error) {
    val, err := getStringResult(32)
    if err != nil {
        return "", err
    }
    if val == "" {
        return NewValue() // pretend function
    }
}

func main() {
    val, err := getVal(32)
    if err != nil {
        panic(err)
    }
    fmt.Println(val)
}
```

### 变量声明

除了避免可变范围和可变性之外,我们还可以提高可读性,但使变量声明接近逻辑。在 C 编程中,通常可以看到以下声明变量的方法:

```go
func main() {
  var err error
  var items []Item
  var sender, receiver chan Item
  
  items = store.GetItems()
  sender = make(chan Item)
  receiver = make(chan Item)
  
  for _, item := range items {
    ...
  }
}
```

其症状与可变范围中所述的症状相同。即使这些变量可能实际上不会在任何时候重新分配,这种风格,将保持读者在他们的脚趾,在所有错误的方式。就像计算机记忆一样,我们的大脑有有限的分配量。必须跟踪哪些变量可能发生突变,以及某些内容是否会更改这些项目,只会使获得代码中所发生情况的良好概述变得更加困难。找出最终返回的值,可能是一场噩梦。因此,为了便于我们的读者(这可能是我们的未来版本),最好声明变量尽可能接近其用法:

```go
func main() {
  var sender chan Item
  sender = make(chan Item)
  
  go func() {
   	for {
    	select {
      case item := <- sender:
        // do something
      }
  	} 
  }()
}
```

但是,通过直接在声明时调用函数,我们可以做得更好。这使得函数逻辑与声明的变量关联起来更加清楚,在前面的示例中,该变量不太清楚。

```go
func main() {
  sender := func() chan Item {
    channel := make(chan Item)
    go func() {
      for {
        select { ... }
      }
    }()
    return channel
  }
}
```

进入整个圆圈后,我们可以移动匿名函数,使其成为命名函数:

```go
func main() {
  sender := NewSenderChannel()
}

func NewSenderChannel() chan Item {
  channel := make(chan Item)
  go func() {
    for {
      select { ... }
    }
  }()
  return channel
}
```

仍然很清楚,我们正在声明一个变量和逻辑,以及与返回的通道关联的逻辑。与第一个示例不同。这使得遍历代码和了解每个变量的责任变得更加容易。

当然,这实际上并限制了我们改变`sender`变量。我们对此无能为力,因为无法在 Go 中声明`const struct`或`static`变量。这意味着,我们必须克制自己,避免在代码的稍后点改变此变量。

> 注: 关键字`const`确实存在,但仅限于在基元类型上使用。

解决这种情况的一种方法,它至少会将变量的可变性限制为包级别。是创建一个结构,将变量作为私有属性。从此以后,这种私有属性只能通过这种包装结构的其他方法访问。扩展我们的频道示例,如下所示:

```go
type Sender struct {
  sender chan Item
}

func NewSender() *Sender {
  return &Sender{
    sender: NewSenderChannel(),
  }
}

func (s *Sender) Send(item Item) {
  s.sender <- item
}
```

我们现在已经确保,我们的`Sender`结构的`sender`属性永远不会发生突变。至少不是,从包外面。在编写本文档时,这是创建公开不可变非原始变量的唯一方法。这有点冗长,但确实值得付出努力,以确保我们不会最终出现奇怪的错误,这可能是结构属性突变的结果。

```go
func main() {
  sender := NewSender()
  sender.Send(&Item{})
}
```

查看上面的示例,可以清楚地了解这如何还简化了包的使用。这种隐藏实现的方式,不仅有利于包的维护者,也有利于包的用户。现在,在初始化和使用`Sender`结构时,无需担心实现。这为更宽松的体系结构打开了一条新天。由于我们的用户不关心实现,我们可以随时更改它,因为我们减少了包的联系人用户。如果我们不再希望在包中使用通道实现,我们可以轻松地更改此方法,而不会中断`Send`方法的使用(只要我们遵循它的当前函数签名)。

> 注: 有一个奇妙的解释,如何处理在客户端库中的抽象,从演讲 AWS re:发明 2017: 拥抱变化而不打破世界 (DEV319)

## Clean Go

本节将介绍编写干净 Go 代码的一些不太通用的方面,而是讨论非常具体的方面。与上一节一样,仍将讨论一些通用和特定的概念,但是,本节标志着文档的开始,其中文档从使用 Go 示例的干净代码的通用描述更改为 Go 特定描述,基于干净的代码原则。

### 返回值

#### 返回已定义错误

我们将开始好一个简单的,通过描述一个更简洁的方式返回错误。与前面讨论的那样,我们编写干净代码的主要目标是确保代码库的可读性、可测试性和可维护性。这种错误返回方法将改进所有三个方面,只需很少的努力。

让我们考虑返回自定义错误的正常方法。这是一个假设示例,取自线程安全映射实现,我们命名为`Store`:

```go
package smelly

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return Item{}, errors.New("item could not be found in the store") 
    }
    return item, nil
}
```

在隔离中,这个功能本身没有什么气味。我们查看`Store`结构`items`地图,看看我们是否已经有一个具有此`id`的项目。如果我们这样做,我们返回项目,如果我们不返回,我们返回错误。相当标准。那么,返回这样的自定义错误有什么问题呢?那么,让我们来看看会发生什么,当我们使用这个函数,从另一个包:

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := smelly.GetItem("123")
    if err != nil {
        if err.Error() == "item could not be found in the store" {
            http.Error(w, err.Error(), http.StatusNotFound)
	        return
        }
        http.Error(w, errr.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

这实际上还不算太糟。然而,有一个明显的问题。Go 中的错误只是实现函数 (`Error()`的`interface`它返回一个字符串。因此,我们现在正在硬编码预期的错误代码到我们的代码库。这不是太好。主要是,因为如果错误消息值更改,我们的代码将中断(柔和)。我们的代码耦合得太紧密,这意味着我们必须在许多不同的位置更改代码。更糟糕的是,如果客户端使用我们的包来编写此代码。如果我们选择更改返回错误的消息,他们的软件会在软件包更新后突然中断。这显然是我们想要避免的。幸运的是,修复非常简单。

```go
package clean

var (
    NullItem = Item{}

    ErrItemNotFound = errors.New("item could not be found in the store") 
)

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, ErrItemNotFound
    }
    return item, nil
}
```

通过将错误转换为变量`ErrItemNotFound`这一简单更改,我们确保使用此包的任何人都可以检查变量,而不是它返回的实际字符串:

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if err == clean.ErrItemNotFound {
           http.Error(w, err.Error(), http.StatusNotFound)
	        return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

感觉好多了,也更安全了。有些人甚至会说,它更容易阅读以及。在出现更详细的错误消息的情况下,开发人员最好只阅读`ErrItemNotFound`而不是一篇关于为什么返回某个错误的小说。

此方法不限于错误,可用于其他返回的值。例如,我们还会像以前一样返回`NullItem`而不是`Item{}`在许多不同的方案中,最好返回定义的对象,而不是在返回时初始化它。

在某些情况下,返回默认`Null`值(如前面的示例)也可能更安全。例如,包的用户可能会忘记检查错误,最终初始化变量,指向一个包含默认值`nil`的空结构作为一个或多个属性值。当稍后尝试在代码中访问此`nil`值时,这可能会导致其代码出现死机。但是,当我们返回自定义默认值时,我们可以确保所有值(否则默认为 nil)都初始化,从而`nil`确保我们不会在用户/客户端软件中造成恐慌。这对我们自己也有好处,就像我们想要实现相同的安全性一样,在不返回默认值的情况下,我们将不得不更改我们的代码,我们返回此类空值的每个位置。但是,使用默认值方法,我们现在只需在一个位置更改代码:

```go
var NullItem = Item{
    itemMap: map[string]Item{},
}
```

> 注意:在许多情况下,调用调用调用实际上更可取。指示缺少错误检查。

> 注: Go 中的每个接口属性的默认值为`nil`。这意味着,对于任何具有接口属性的结构,这都很有用。对于包含通道、地图和切片的结构也是如此,它们可能也具有`nil`值。

#### 返回动态错误

当然,在某些情况下,返回错误变量实际上可能不可行。在自定义错误信息是动态的的情况下,为了更具体地描述错误事件,我们不能再定义和返回静态错误。例如:

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, fmt.Errorf("Could not find item with ID: %s", id)
    }
    return item, nil
}
```

那么,该怎么办呢?没有定义良好的/标准方法来处理和返回这些类型的动态错误。我个人的喜好是返回一个新的界面,并附加了一些功能:

```go
type ErrorDetails interface {
    Error() string
    Type() string
}

type errDetails struct {
    errtype error
    details string
}

func NewErrorDetails(err error, details ...interface{}) ErrorDetails {
    return &errDetails{
        errtype: err,
        details: details,
    }
}

func (err *errDetails) Error() string {
    return fmt.Sprintf("%v: %v", err.details)
}

func (err *errDetails) Type() error {
    return err.errtype
}
```

这种新的数据结构仍然作为我们的标准错误。我们仍然可以将其与`nil`进行比较,因为它是接口实现,我们仍然可以调用`.Error()`上,所以它不会破坏任何现有的实现。但是,好处是,我们现在可以像以前一样检查错误类型,尽管我们的错误现在包含*动态*详细信息:

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, fmt.Errorf("Could not find item with ID: %s", id)
    }
    return item, nil
}
```

然后,可以重构我们的 http 处理程序函数,以再次检查特定错误:

```go
func GetItemHandler(w http.ReponseWriter, r http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if err.Type() == clean.ErrItemNotFound {
            http.Error(w, err.Error(), http.StatusNotFound)
	        return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

### 零值

围棋的一个有争议的方面是增加`nil`此值对应于 C 中`NULL`的值 NULL,本质上是一个未初始化的指针。以前,我们在`nil`可以引起的麻烦中进行了介绍,但可以总结一下:当您尝试访问方法或属性为`nil`值时,事情会中断。在上述部分中,建议尽量减少返回`nil`值的使用。这样,我们的代码用户就不太容易因错误而意外访问`nil`值。

在其他情况下,查找`nil`值很常见,这可能会导致一些不必要的痛苦。例如,`struct`的初始化不正确会导致`struct`包含`nil`属性。如果访问,它们将导致恐慌。下面可以看到这方面的一个例子:

```go
type App struct {
	Cache *KVCache
}

type KVCache struct {
  mtx sync.RWMutex
	store map[string]string
}

func (cache *KVCache) Add(key, value string) {
  cache.mtx.Lock()
  defer cache.mtx.Unlock()
  
	cache.store[key] = value
}
```

此代码绝对正常。但是,我们的`App`可以错误地初始化,而无需在其中初始化我们的`Cache`属性,这一事实暴露了我们。如果调用以下代码,我们的应用程序将死机:

```go
	app := App{}
	app.Cache.Add("panic", "now")
```

`Cache`属性从未初始化,因此是`Add`方法的`nil`指针。运行此代码将导致死机,并显示以下消息:

> 死机:运行时错误:无效的内存地址或 0 指针取消引用

相反,我们可以将`App`结构的`Cache`属性转换为私有属性,并创建类似于 getter 的方法,以访问`App`的`Cache`属性。这让我们可以更控制我们返回的内容,并确保我们不会返回`nil`值。

```go
type App struct {
    cache *KVCache
}

func (app *App) Cache() *KVCache {
	if app.cache == nil {
    app.cache = NewKVCache()
	}
	return app.cache
}
```

现在,在尝试访问`Cache`属性时,我们确保永远不会遇到返回`nil`指针。以前惊慌失措的代码现在将重构为以下内容:

```go
app := App{}
app.Cache().Add("panic", "now")
```

这样做更可取的原因是,我们确保包的用户不会担心实现,以及他们是否以不安全的方式使用我们的包。他们需要担心的只是编写自己的干净代码。

> 注意:还有其他方法可以实现类似的安全结果,但是,我认为这是最直接的方法。

### 转到中的指针

指针是一个很大的话题。他们是使用语言的一个非常重要的部分,以至于在不了解指针和它们的工作的情况下,基本上不可能写去。在本文中,我将不详细介绍指针的内部工作。相反,我们将专注于他们的怪癖,以及如何处理他们去。

但是,指针增加了复杂性,但是,正如前面提到的,在编写时几乎不可能避免它们。因此,了解如何使用指针非常重要,同时不增加不必要的复杂性,从而保持代码库的整洁。在不限制自己的情况下,不正确使用指针可能会带来令人讨厌的副作用,引入特别难以调试的 Bug。当然,当坚持编写干净代码的基本原则时,本文的第一部分介绍了,我们限制了引入这种复杂性的曝光,但指针是一个特例,它仍然可以撤消我们以前的所有辛勤工作,使我们的代码干净。

#### 指针可变性

我在本文中已经不止一次地使用了"易变性"这个词,作为否定词。可变性显然不是一个明确的坏事,我绝不是一个倡导者,写100%纯功能程序。可变性是一个强大的工具,但我们真的应该只使用它,当它是必要的。让我们看一下代码示例,说明原因:

```go
func (store *UserStore) Insert(user *User) error {
    if store.userExists(user.ID) {
        return ErrItemAlreaydExists
    }
    store.users[user.ID] = user
    return nil
}

func (store *UserStore) userExists(id int64) bool {
    _, ok := store.users[id]
    return ok
}
```

乍一看,这看起来还不算太糟。事实上,它甚至看起来像一个相当简单的插入函数,用于公共列表结构。我们接受指针作为输入,如果不存在具有此`id`的其他用户,则我们将用户指针插入到我们的列表中。现在,我们在公共 API 中使用此功能来创建新用户:

```go
func CreateUser(w http.ResponseWriter, r *http.Request) {
    user, err := parseUserFromRequest(r)
    if err != nil {
        http.Error(w, err, http.StatusBadRequest)
        return
    }
    if err := insertUser(w, user); err != nil {
      http.Error(w, err, http.StatusInternalServerError)
      return
    }
}

func insertUser(w http.ResponseWriter, user User) error {
  	if err := store.Insert(user); err != nil {
        return err
    }
  	user.Password = ""
	  return json.NewEncoder(w).Encode(user)
}
```

再一次,乍一看一切看起来都很好。我们从收到的请求中分析用户,并将用户结构插入到我们的商店中。一旦我们成功地将用户插入存储中,我们将密码设置为零,然后再将用户作为 JSON 对象返回给客户端。这是所有相当常见的做法,通常当返回用户对象时,如果密码已进行哈希,则我们不想返回哈希密码。

但是,假设我们使用的是基于`map`的内存存储,此代码将产生一些意外的结果。如果我们签入用户存储,则看到我们在 http 处理程序函数中对用户密码所做的更改也影响了我们存储中的对象。这是因为`parseUserFromRequest`返回的指针地址是我们填充存储时使用的,而不是实际值。因此,在更改取消引用的密码值时,我们最终将更改存储中指向的对象的值。

这是一个很好的例子,说明为什么可变性和可变范围都会导致一些严重的问题和错误,当使用不正确。当将指针作为函数的输入参数传递时,我们将扩展变量的范围。更令人担忧的是,我们正在将范围扩大到未界定的水平。我们*几乎*将变量的范围扩展到全局可用的变量。取决于我们商店的可变范围。如上例所示,这可能导致灾难性的错误,这些错误尤其难以发现和根除。

幸运的是,这个Bug的修复程序相当简单:

```go
func (store *UserStore) Insert(user User) error {
    if store.userExists(user.ID) {
        return ErrItemAlreaydExists
    }
    store.users[user.ID] = &user
    return nil
}
```

我们现在传递的是`User`结构的副本,而不是将指针传递给`User`结构。但是,我们仍存储指向存储的指针,而不是从函数外部存储指针,而是存储指向复制值的指针,该值位于函数内。这修复了眼前的问题,但如果我们不小心的话,可能还会进一步导致问题。

```go
func (store *UserStore) Get(id int64) (*User, error) {
    user, ok := store.users[id]
    if !ok {
        return EmptyUser, ErrUserNotFound
    }
    return store.users[id], nil
}
```

同样,我们商店的 getter 函数的非常标准的非常简单的实现。然而,这仍然是不好的。我们再次扩展了指针的范围,这可能会导致意外的副作用。当返回实际指针值(我们存储在用户存储中)时,我们实质上是使应用程序的其他部分能够更改存储值。这是不好的,因为它必然要确保混乱。我们的商店应该是唯一启用对存储的值进行更改的实体。最简单的解决方法是返回`User`的值,而不是返回指针。

> 注意:我们的应用程序应该使用多个线程,这种情况通常是这样。将指针传递到同一内存位置,也可能导致争用条件。换句话说,我们不仅可能损坏数据,还可能导致数据竞赛的恐慌。

请记住,返回指针在本质上没有错,但是变量的扩展范围和所有者数量是重要的方面。这就是我们前面的示例对有臭味的操作进行分类的原因。这也是为什么,常见的Go构造函数也绝对罚款:

```go
func AddName(user *User, name string) {
    user.Name = name
}
```

这样做的原因,是变量作用域在函数返回后保持不变,该范围由调用函数的哪个人定义。这与变量的所有权保持不变(仅与函数调用器保持一起)的面相结合,意味着无法以意外的方式操作指针。

### 闭包是函数指针

因此,在转到在 Go 中使用接口的下一个主题之前。我想介绍一下通常监督的替代方法,即 C 程序员所说的"函数指针",而大多数其他程序员称之为"闭合"。闭合非常简单。它们是函数的输入参数,它的作用与任何其他参数类似,但它们是函数。在 Javascript 中,使用闭包作为回调很常见,这通常用于在异步操作完成后调用函数的情况中。在Go中,我们并没有真正解决这个问题,或者至少,我们有其他更好的方法来解决这个问题。相反,在 Go 中,我们可以使用闭包来解决另一个障碍:缺少泛型。

现在,不要太激动。我们不会取代泛型的缺乏。我们只是用闭包来解决泛型缺乏的子集。请考虑以下函数签名:

```go
func something(closure func(float64) float64) float64 { ... }
```

此函数以另一个函数作为输入,并将返回`float64`输入函数,将`float64`作为输入,也将返回`float64`此模式对于创建松散耦合的体系结构特别有用,可以更轻松地添加功能,而不影响代码的其他部分。这方面的一个示例用例可能是包含数据结构,我们希望以某种形式操作这些数据。通过此结构`Do()`方法,我们可以对此数据执行操作。如果我们提前知道操作,我们可以通过将处理不同操作的逻辑直接放在`Do()`方法中来处理问题:

```go
func (datastore *Datastore) Do(operation Operation, data []byte) error {
  switch(operation) {
  case COMPARE:
    return datastore.compare(data)
  case CONCAT:
    return datastore.add(data)
  default:
    return ErrUnknownOperation
  }
}
```

可以想象,此函数将对`Datastore`结构中包含的数据执行预定操作。但是,我们也可以想象,在某些时候,我们希望添加更多操作。在较长时间内,这可能会成为许多不同的操作,使我们的`Do`方法膨胀,甚至可能难以维护。对于想要使用我们的`Datastore`对象的用户来说,这也可能成为一个问题,他们无法访问我们的包代码。请记住,无法像大多数 OOP 语言那样扩展结构方法。对于想要使用我们的软件包的开发人员来说,这也可能成为一个问题。

因此,让我们尝试一种不同的方法,改为使用闭包:

```go
func (datastore *Datastore) Do(operation func(data []byte, data []byte) ([]byte, error), data []byte) error {
  result, err := operation(datastore.data, data)
  if err != nil {
    return err
  }
  datastore.data = result
  return nil
}

func concat(a []byte, b []byte) ([]byte, error) {
  ...
}

func main() {
  ...
  datastore.Do(concat, data)
  ...
}
```

但是,除了这是一个非常混乱的函数签名之外,我们还有另一个问题。此函数不是特别通用。如果我们发现实际上希望`concat`函数需要能够将多个字节数组作为输入,会发生什么情况?或者,如果要添加一些全新的功能,可能还需要更多或更少的输入`(data []byte, data []byte)` ?

解决这个问题的一个方法是改变我们的concat函数。在下面的示例中,我将其更改为仅将单个字节数组作为输入参数,但情况可能正好相反。

```go
func concat(data []byte) func(data []byte) ([]byte, error) {
  return func(concatting []byte) ([]byte, error) {
    return append(data, concatting), nil
  }
}

func (datastore *Datastore) Do(operation func(data []byte) ([]byte, error)) error {
  result, err := operation(datastore.data)
  if err != nil {
    return err
  }
  datastore.data = result
  return nil
}

func main() {
  ...
  datastore.Do(compare(data))
  ...
}
```

请注意,我们如何从`Do`方法签名中添加一些杂乱。我们实现此目的的方式是让`concat`函数返回函数。在返回的函数中,我们将存储最初传入到`concat`函数的输入值。因此,返回的函数现在可以采用单个输入参数,在我们的函数逻辑中,我们将用原始输入值追加它。作为一个新引入的概念,这是很奇怪的,但是,习惯有这个作为一个选项确实可以帮助放松程序耦合,并帮助摆脱膨胀的功能。

在下一节中,我们将讨论接口,但让我们花一小会儿来讨论接口和闭包之间的区别。接口解决的问题肯定与闭包解决的问题重叠。在 Go 中实现接口可以区分何时使用其中一个接口,有时有些困难。通常,无论使用接口还是闭包,都不太重要,无论以最简单的方式解决问题,都是正确的选择。通常,如果操作本质上很简单,则关闭将更容易实现。但是,一旦闭包中包含的逻辑变得复杂,就应该强烈建议使用接口。

戴夫·切尼对这个话题有一个很好的写作,并谈了同一个话题:

- [https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions](https://www.translatoruser-int.com/translate?&to=zh-CHS&csId=24f5af06-cfe4-44a8-9412-da1a7e7ea8c9&usId=7f43ce80-5598-418d-ac38-b19975567286&dl=en&ac=true&dt=2019%2f6%2f23 10%3a19&h=Xgo1ZxWPAHYM0lDgymxhpd_CZZIDIw63&a=https%3a%2f%2fdave.cheney.net%2f2016%2f11%2f13%2fdo-not-fear-first-class-functions)
- [https://www.youtube.com/watch?v=5buaPyJ0XeQ&t=9s](https://www.translatoruser-int.com/translate?&to=zh-CHS&csId=24f5af06-cfe4-44a8-9412-da1a7e7ea8c9&usId=7f43ce80-5598-418d-ac38-b19975567286&dl=en&ac=true&dt=2019%2f6%2f23 10%3a19&h=HedsZGuoqrE05BUugk6MP5KNBpk7-FIZ&a=https%3a%2f%2fwww.youtube.com%2fwatch%3fv%3d5buaPyJ0XeQ%26amp%3bt%3d9s)

乔恩·博德纳也谈到了这个话题

- [https://www.youtube.com/watch?v=5IKcPMJXkKs](https://www.translatoruser-int.com/translate?&to=zh-CHS&csId=24f5af06-cfe4-44a8-9412-da1a7e7ea8c9&usId=7f43ce80-5598-418d-ac38-b19975567286&dl=en&ac=true&dt=2019%2f6%2f23 10%3a19&h=JlqcBcinW8-bVdF0_ZJlyHChDZkN7T5F&a=https%3a%2f%2fwww.youtube.com%2fwatch%3fv%3d5IKcPMJXkKs)

### 转到中的接口

通常,处理`interface`的 go 方法与其他语言有很大不同。接口没有显式实现,就像它们在 Java 或 C# 中一样,但如果它们实现了接口的协定,则隐式实现。例如,这意味着任何具有`Error()`方法的`struct`实现/实现`Error`接口,并可以返回为`error`。这有它的优点,因为它让Go感觉更快节奏和动态,因为接口实现是非常容易的。实现接口的这种方法显然也有缺点。由于接口实现不再显式,因此很难看到哪些接口由结构实现。因此,定义接口的最常见方法是用尽可能少的方法编写接口。这样,就更容易理解结构是否实现了接口的协定。

还有其他方法来跟踪结构是否正在履行接口协定。一种方法是创建返回接口的构造函数,而不是具体类型:

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}

type NullWriter struct {}

func (writer *NullWriter) Write(data []byte) (n int, err error) {
    // do nothing
    return len(data), nil
}

func NewNullWriter() io.Writer {
    return &NullWriter{}
}
```

上述函数确保`NullWriter`结构实现`Writer`接口。如果我们删除`NullWriter`的`Write`方法,我们将收到编译错误,如果我们尝试构建解决方案。这是一种确保代码按预期方式的行为的好方法,并且我们可以将编译器用作安全网,以确保我们不会生成无效代码。

还有一种方法可以更明确地了解给定结构实现的接口。然而,这种方法实现了与我们希望实现的结果相反的效果。方法使用嵌入接口作为结构属性。

"等什么?" -大概大多数人

所以,让我们倒一点,然后潜入臭气熏天的禁林。在 Go 中,我们可以使用嵌入式结构,作为结构定义中的继承类型。这真的很好,因为我们可以通过定义可重用结构来分离代码。

```go
type Metadata struct {
    CreatedBy types.User
}

type Document struct {
    *Metadata
    Title string
    Body string
}

type AudioFile struct {
    *Metadata
    Title string
    Body string
}
```

在上面,我们正在定义一个`Metadata`对象,它将为我们提供属性字段,我们可能用于许多不同的结构类型。使用嵌入式结构(而不是直接在我们的结构中显式定义属性)的妙处是,它已分离了`Metadata`字段。应选择更新元数据对象,`Metadata`我们可以在单个位置进行更改。如前所述,我们希望确保代码中的一个位置的更改不会破坏代码的其他部分。保持这些属性的集中化,将使用户清楚使用嵌入`Metadata`的结构具有相同的属性。就像,满足接口的结构具有相同的方法。

现在,让我们看一个示例,说明在更改`Metadata`结构时,如何使用构造函数以进一步防止破坏我们的代码:

```go
func NewMetadata(user types.User) Metadata {
    return &Metadata{
        CreatedBy: user,
    }
}

func NewDocument(title string, body string) Document {
    return Document{
        Metadata: NewMetadata(),
        Title: title,
        Body: body,
    }
}
```

稍后,我们发现,我们还需要`Metadata`对象上的`CreatedAt`字段。现在,只需更新`NewMetadata`构造函数,即可轻松实现:

```go
func NewMetadata(user types.User) Metadata {
    return &Metadata{
        CreatedBy: user,
        CreatedAt: time.Now(),
    }
}
```

现在,我们的`Document`和`AudioFile`结构都进行了更新,以便在构造时填充这些字段。这是脱钩背后的核心原则,也是确保代码可维护性的绝佳示例。我们还可以添加新方法,而不会破坏我们的代码:

```go
type Metadata struct {
    CreatedBy types.User
    CreatedAt time.Time
    UpdatedBy types.User
    UpdatedAt time.Time
}

func (metadata *Metadata) AddUpdateInfo(user types.User) {
    metadata.UpdatedBy = user
    metadata.UpdatedAt = time.Now()
}
```

同样,在不破坏代码库的其余部分的情况下,我们正在实现现有结构的新功能。这种编程使得实现新功能非常快速且非常无痛,这正是我们试图通过清理代码来实现的目标。

现在,我很抱歉打破这种幸福的条纹,因为现在我们回到了臭气熏天的禁火森林。让我们回到我们的接口以及如何显式显示哪些接口正在由结构实现。我们可以嵌入接口,而不是嵌入结构:

```
type NullWriter struct {
    Writer
}

func NewNullWriter() io.Writer {
    return &NullWriter{}
}
```

上述代码进行编译。我第一次看到这个,我简直不敢相信,这实际上是在编译。从技术上讲,我们正在实现Writer的接口,因为我们`Writer`正在嵌入接口并"继承"与此接口关联的函数。有些人认为这是一种明确的方式,表明我们的`NullWriter`正在实现`Writer`接口。但是,我们必须谨慎使用此技术,因为我们不能再依赖编译器来保存我们:

```
func main() {
    w := NewNullWriter()

    w.Write([]byte{1, 2, 3})
}
```

如前所述,上述代码将编译。根据`NewNullWriter`返回一个`Writer`一切都是 - dori , 因为`NullWriter`履行了`io.Writer`, 通过.嵌入式接口。但是,运行上述代码将导致以下情况:

> 死机:运行时错误:无效的内存地址或 0 指针取消引用

解释是,在Go中的接口方法,本质上是一个函数指针。在这种情况下,由于我们指向的是接口的函数,而不是实际的方法实现,所以我们尝试调用一个函数,实际上它是一个 nil 指针。哎呀！就个人而言,我认为这是 Go 编译器中的一个巨大疏忽。此代码**不应**编译...但是,当这是正在修复(如果它将永远是),让我们只是承诺对方,从来没有以这种方式实现代码。为了更清楚地说明我们的实现,我们最终还是自己在脚下拍摄了,并绕过了编译器检查。

> 有些人认为,使用嵌入式接口是创建模拟结构的好方法,用于测试接口方法的子集。从本质上讲,通过使用嵌入式接口,您不必实现接口的所有方法,而只是实现希望测试的少数方法。在测试/嘲弄中,我可以看到论点,但我仍然不是这种方法的粉丝。

让我们快速返回清理代码,并快速返回在 Go 中正确使用接口。让我们来讨论使用接口作为函数参数和返回值。在 Go 中使用函数的接口使用最常见的谚语是:

"在做人时要保守,在别人接受的东西上要自由" - 乔恩·波斯特尔

> 事实:这句谚语最初与Go无关,但实际上取自TCP网络协议的早期规范。

换句话说,您应该编写接受接口并返回具体类型的函数。这通常是好的做法,在用嘲弄做测试时变得超级有益。例如,我们可以创建一个函数,该函数将编写器接口作为输入并调用该接口的`Write`方法。

```go
type Pipe struct {
    writer io.Writer
    buffer bytes.Buffer
}

func NewPipe(w io.Writer) *Pipe {
    return &Pipe{
        writer: w,
    }
} 

func (pipe *Pipe) Save() error {
    if _, err := pipe.writer.Write(pipe.FlushBuffer()); err != nil {
        return err
    }
    return nil
}
```

假设我们在应用程序运行时写入文件,但我们不希望为调用此函数的所有测试写入新文件。因此,我们可以实现一个新的模拟类型,这基本上不会执行任何操作。从本质上讲,这只是基本的依赖项注入和模拟,但问题是,它是非常容易使用在去:

```go
type NullWriter struct {}

func (w *NullWriter) Write(data []byte) (int, error) {
    return len(data), nil
}

func TestFn(t *testing.T) {
    ...
    pipe := NewPipe(NullWriter{})
    ...
}
```

> 注意: 实际上已经有一个内置于`Discard`的 ioutil 包中的空编写器实现

使用`NullWriter`而不是其他编写器)构造管道结构时,在调用`Save`函数时,不会发生任何操作。 `Pipe`我们唯一要做的就是添加4行代码。这就是为什么在习惯性去,它被鼓励使接口类型尽可能小,使实现这样的模式尽可能容易。然而,这种接口的实现,也带来了巨大的缺点。

### 空接口

与其他语言不同,go 没有泛型的实现。已经提出了许多实施建议,但Go语言小组认为所有这些建议都不满意。遗憾的是,如果没有泛型,开发人员会尝试找到解决此问题的创造性方法,经常使用空`interface{}`下一节将介绍为什么这些通常过于创造性的实现应该被视为不良做法和不干净的代码。还将有空`interface{}`的使用示例* 以及如何避免使用空`interface{}`编写代码的一些陷阱。.

但首先也是最重要的。是什么促使开发人员使用空`interface{}`?嗯,正如我前面说的,Go确定具体类型是否实现接口的方式是检查它是否实现了特定接口的方法。那么,如果我们的接口根本不实现任何方法,会发生什么情况呢?

```
type EmptyInterface interface {}
```

上述等效于内置类型`interface{}`。此接口类型的结果是接受**任何**类型。这意味着,我们可以编写接受任何类型的函数。这对于某些类型的功能(如创建打印机功能时)非常有用。这是如何从`fmt`包为`Println`函数提供任何类型:

```
func Println(v ...interface{}) {
    ...
}
```

在这种情况下,我们不仅接受单个`interface{}`而且接受一个类型切片。这些类型可以是任何类型,甚至可以是不同类型的,只要它们实现空`interface{}`我们确信任何类型都会实现。这是处理字符串会话(从字符串到字符串)时非常常见的模式。原因是,这是 Go 中实现泛型方法的唯一方法。这方面的好例子来自`json`标准库包:

```go
func InsertItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.InsertItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatsOK)
}
```

所有*不太优雅的*代码都包含在"解码"函数`Decode`中。因此,使用此功能的开发人员不必担心类型的反射或强制转换。我们只需要担心提供指向具体类型的指针。这很好,因为`Decode()`函数在技术上返回了具体类型。我们正在传递`Item`值,该值将从 http 请求的正文填充,我们不必处理处理`interface{}`值的潜在风险。

但是,即使使用具有良好做法的空`interface{}`我们仍有一些问题。如果我们传递的 JSON 字符串与我们的`Item`类型无关,但仍有效,我们仍不会收到错误。我们的`item`变量将只保留默认值。因此,虽然我们不必担心反射和强制转换错误,但我们仍必须确保从客户端发送的消息是有效的`Item`类型。但是,在编写本文档时,没有简单/好的方法来实现这些类型的泛型解码器,而无需使用空`interface{}`类型。

问题在于,我们倾向于使用 Go(一种静态类型化语言)作为动态类型化语言。当查看`interface{}`糟糕的实现时,这一点变得更加清晰。最常见的示例来自尝试实现某种通用存储/列表的开发人员。让我们看一个示例,尝试实现一个通用的 HashMap 包,该包可以使用`interface{}`存储任何类型。.

```go
type HashMap struct {
    store map[string]interface{}
}

func (hashmap *HashMap) Insert(key string, value interface{}) {
    hashmap.store[key] = value
}

func (hashmap *HashMap) Get(id string) (interface{}, error) {
    value, ok := hashmap.store[key]
    if !ok {
        return nil, ErrKeyNotFoundInHashMap
    }
    return value
}
```

> 注意:为了简单起见,我省略了示例中的线程安全性

请记住,上面使用的实现模式在相当多的 Go 包中使用。它甚至用于标准库`sync`包,用于`sync.Map`类型。那么,这个实现有什么大问题呢?好吧,让我们来看看使用此包的示例。

```go
func SomeFunction(id string) (Item, error) {
    itemIface, err := hashmap.Get(id)
    if err != nil {
        return EmptyItem, err
    }
    item, ok := itemIface.(Item)
    if !ok {
        return EmptyItem, ErrCastingItem
    }
    return item, nil
}
```

乍一看,这看起来很好。但是,如前所述。然而,如果我们在我们的商店中添加不同类型的产品,我们将开始陷入麻烦,而到目前为止,这些类型并没有被阻止。除了`Item`类型之外,没有任何限制我们添加其他内容。那么,当有人开始向我们的哈希地图添加其他类型的内容时,会发生什么情况呢?我们的函数现在可能会返回错误。这甚至可能是一个小的变化,就像代码库中的其他人想要存储指针`*Item`而不是`Item`。最糟糕的是,这甚至可能不是我们的测试所捕捉的。根据系统的复杂性,这可能会引入一些特别难以调试的 Bug。

这种类型的代码,不应该达到生产。事实是,Go现在不支持泛型,作为Go程序员,我们应该接受这一点。如果我们想要使用泛型,那么我们应该使用支持泛型的不同语言,而不是试图破解我们的方式。

那么,我们如何防止此代码进入生产?我们问题的简单解决方案,基本上是只使用具体类型的函数编写,而不是使用`interface{}`值。当然,这并不总是最好的方法,因为包中可能有一些实现自己的功能。因此,创建包装器可能是一种更好的方法,它公开了我们需要的功能,但仍确保类型安全:

```go
type ItemCache struct {
  kv tinykv.KV
} 

func (cache *ItemCache) Get(id string) (Item, error) {
  value, ok := cache.kv.Get(id)
  if !ok {
    return EmptyItem, ErrItemNotFound
  }
  return interfaceToItem(value)
}

func interfaceToItem(v interface{}) (Item, error) {
  item, ok := v.(Item)
  if !ok {
    return EmptyItem, ErrCouldNotCastItem
  }
  return item, nil
}

func (cache *ItemCache) Put(id string, item Item) error {
  return cache.kv.Put(id, item)
}
```

> 注:小kv的其他功能的实现。为了简洁起见,KV 缓存被遗漏了。

现在,创建上面的包装器将确保我们使用实际类型,并且不再传入`interface{}`类型。因此,我们不再可能意外地用错误的值类型填充我们的商店,并且我们已经尽可能包含了类型的强制转换。这是解决我们问题的一个非常直截了当的方法,虽然有点手动。

## 总结

首先,感谢您通过本文一路走来。我希望它能够深入了解什么是干净代码,以及它将如何帮助确保代码库中的可维护性、可读性和稳定性。总结涵盖的所有主题:

**函数**- 函数的命名应该变得更加具体,函数的范围越小。确保所有功能都是单一用途。一个很好的度量,是限制您的函数长度到 5-8 行,并且只需要 2-3 个输入参数。

**变量**- 变量的命名应该越不具体,范围越小,并且将变量的范围保持在最小。此外,将变量的可变性降至最低,并随着变量范围的增长,越来越意识到这一点。

**返回值**- 应尽可能返回具体类型。尽可能使包的用户难以产生错误,并让他们易于理解函数返回的值

**指针**- 谨慎使用指针,并将范围和可变性限制为绝对最小值。垃圾回收仅有助于内存管理,它无助于处理与指针关联的所有其他复杂性。

**接口**- 尽可能使用接口来放松代码的耦合。尽可能包含使用空接口的任何`interface{}`并防止其公开。

当然,什么被认为是干净的代码是特别主观的,我不认为这永远不会改变。然而,就像我`gofmt`的陈述一样,我认为找到一个共同的标准比每个人都同意100%的标准更重要。理解狂热从来不是目标也很重要。代码库很可能永远不会是 100%"干净"的,就像你的办公桌不是一样。有超越本文中确立的规则和边界的空间。但是,请记住,编写干净代码的最重要方面是互相帮助。我们通过确保软件的稳定性和易于调试来帮助我们的支持工程师。我们通过确保我们的代码可读且易于消化来帮助其他开发人员。我们通过建立灵活的代码库来帮助参与项目的每个人,在代码库中,我们可以快速引入新功能,而不会破坏我们当前的平台。我们走得很慢,从此,每个人都很满意。

因此,我希望你们能加入讨论,帮助我们,Go社区,定义为干净的代码。让我们建立一个共同点,以便我们改进软件。不仅为了我们自己,也为了每个人。

