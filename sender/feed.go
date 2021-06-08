package sender

import (
	"fmt"
	"io/ioutil"
	"strings"

	aggregator "github.com/wwhtrbbtt/PersonalNewsletter/aggregator"
)

type Feed struct {
	Time         string              `json:"time"`         // ex. "10:00"
	Feedname     string              `json:"feedName"`     // ex. "My email feed"
	Greetingname string              `json:"greetingName"` // ex. "peet"
	Email        string              `json:"email"`        // ex. peet@peet.ws
	Modules      []aggregator.Module `json:"modules"`      // collection of modules
}

func (f Feed) GetHTML(template string) string {
	// the template has 4 {{things}} to replace:
	// - {{Title}}
	// - {{Greeting}}
	// - {{Modules}}
	// - {{Footer}}

	// Building the module HTML

	MODULES := ""
	for _, module := range f.Modules {
		tmpHTML := ""
		// module title
		tmpHTML += GetTitleHTML(module.Caption)

		// module text
		if module.Text != "" {
			tmpHTML += GetTextHTML(module.Text)
		}
		// Module Button
		if module.Button.Text != "" && module.Button.URL != "" {
			tmpHTML += GetButtonHTML(module.Button)
		}
		// TODO:
		if len(module.Links) > 0 {
			tmpHTML += GetLinksHTML(module.Links)
		}
		// module.Image

		MODULES += tmpHTML
	}

	TITLE := f.Feedname + " - PersonalNewsletter"
	GREETING := fmt.Sprintf("Hello %s!", f.Greetingname)
	FOOTER := "PersonalNewsletter"

	content, err := ioutil.ReadFile(template)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	HTML := string(content)

	// Replace things
	// I know that html/template would be a better choice, but I cant wrap my head around it

	HTML = strings.ReplaceAll(HTML, "{{Title}}", TITLE)
	HTML = strings.ReplaceAll(HTML, "{{Greeting}}", GREETING)
	HTML = strings.ReplaceAll(HTML, "{{Modules}}", MODULES)
	HTML = strings.ReplaceAll(HTML, "{{Footer}}", FOOTER)
	return HTML
}

// get HTML for different elements
func GetButtonHTML(b aggregator.Button) string {
	return fmt.Sprintf(`
	<table role="presentation" border="0" cellpadding="0" cellspacing="0" class="btn btn-primary">
		<tbody>
			<tr>
				<td align="left">
					<table role="presentation" border="0" cellpadding="0" cellspacing="0">
						<tbody>
							<tr>
								<td> <a href="%s" target="_blank">%s</a> </td>
							</tr>
						</tbody>
					</table>
				</td>
			</tr>
		</tbody>
	</table><br>
	`, b.URL, b.Text)
}

func GetTitleHTML(title string) string {
	return fmt.Sprintf(`
	<p><b>%s</b></p>
	`, title)
}
func GetTextHTML(text string) string {
	return fmt.Sprintf(`
	<p>%s</p>
	`, strings.ReplaceAll(text, "\n", "<br>"))
}

func GetLinksHTML(links []aggregator.Link) string {
	html := ""
	for _, link := range links {
		html += fmt.Sprintf(`<a href="%s" target="_blank">%s</a><br>`, link.URL, link.Text)
	}
	return html
}
