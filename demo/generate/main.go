package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	// gqlgen configuration
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(1)
	}

	err = api.Generate(
		cfg,
		api.AddPlugin(enableAllResolvers{}),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

type enableAllResolvers struct{}

func (e enableAllResolvers) Name() string {
	return "enable_all_resolvers"
}

func (e enableAllResolvers) MutateConfig(c *config.Config) error {
	for _, t := range c.Schema.Types {
		// fmt.Printf("type %s\n", t.Name)
		// fmt.Printf(" - IsInputType %t\n", t.IsInputType())
		// fmt.Printf(" - IsLeafType %t\n", t.IsLeafType())
		// fmt.Printf(" - IsCompositeType %t\n", t.IsCompositeType())
		// fmt.Printf(" - IsAbstractType %t\n", t.IsAbstractType())

		for _, f := range t.Fields {
			// fmt.Printf(" - Field %s:\n", f.Name)
			// fmt.Printf("     Name: %s\n", f.Name)
			// fmt.Printf("     Type: %s\n", f.Type)

			typ := f.Type.String()
			// fmt.Printf("     typ is %s\n", typ)
			typ = strings.ReplaceAll(typ, "[", "")
			typ = strings.ReplaceAll(typ, "]", "")
			typ = strings.ReplaceAll(typ, "!", "")
			// fmt.Printf("     typ after is %s\n", typ)

			if ft, ok := c.Schema.Types[typ]; ok {
				if ft.IsCompositeType() {
					fmt.Printf("%s.%s needs a resolver\n", t.Name, f.Name)
				}
			}
		}

	}

	t := c.Models["User"]

	t.Fields = map[string]config.TypeMapField{
		"articles": {
			FieldName: "articles",
			Resolver:  true,
		},
	}

	c.Models["User"] = t

	for i, m := range c.Models {
		// if i == "User" {
		fmt.Printf("model %s: %+v\n", i, m)
		// }
	}

	return nil
}

func (e enableAllResolvers) GenerateCode(d *codegen.Data) error {
	for i, o := range d.Objects {
		if o.Name == "User" {
			fmt.Printf("object %d: %+v\n", i, o)

			// for j, f := range o.Fields {
			// 	fmt.Printf(" - field %s (%d): %+v\n", f.Name, j, f)

			// 	if f.Name == "comments" {
			// 		d.Objects[i].Fields[j].IsResolver = true
			// 	}
			// }
		}
	}

	return nil
}
