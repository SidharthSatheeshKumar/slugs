package slugs

import "strings"

func SlugGenerator(value string) string {
    value = strings.TrimSpace(value)
    slug := strings.ToLower(value)
    slug = strings.ReplaceAll(slug, " ", "-")
    return slug
}