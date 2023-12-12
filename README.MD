# Style your log

## Import the Package:
Import the chalkGo package into your Go code by using the import statement like this:

``` go
import "github.com/bobiscool/chalkGo"
```
## Create a ChalkGo Instance:
To style your text, create an instance of the ChalkGo struct like this:

``` go
chalk := chalkGo.ChalkGo{}
```
## Apply Text Styles:
You can apply various text styles and colors to your text using the methods provided by the ChalkGo instance. For example:

To make the text red, use chalk.Red().
To make the text bold, use chalk.Bold().
To make the text underlined, use chalk.Underline().
You can chain these methods to apply multiple styles to the same text.

## Convert to String:
After applying the desired styles, convert your text to a stylized string using the ToString method:

``` go
stylizedText := chalk.Red().Bold().Underline().ToString("This is a red, bold, and underlined text")
```
## Print or Use the Stylized Text:
You can now use the stylizedText variable in your code. For example, you can print it to the console:

``` go
fmt.Println(stylizedText)
```
This will display the stylized text in the console with the specified styles.

Remember to replace the example text and styles with your own content as needed.