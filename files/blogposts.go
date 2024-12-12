package blogposts

import (
  "bufio"
  "bytes"
  "fmt"
  "io"
  "io/fs"
  "strings"
)

type Post struct {
  Title, Description, Body string
  Tags []string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
  dir, err := fs.ReadDir(fileSystem, ".")
  if err != nil {
    return nil, err
  }

  var posts []Post
  for _, f := range dir {
    post, err := getPost(fileSystem, f)
    if err != nil {
      return nil, err
    }
    posts = append(posts, post)
  }
  return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
  postFile, err := fileSystem.Open(f.Name())
  if err != nil {
    return Post{}, err
  }
  defer postFile.Close()

  return newPost(postFile)
}

const (
  titleSeparator        = "Title: "
  descriptionSeparator  = "Description: "
  bodySeparator         = "Body: "
  tagsSeparator         = "Tags: "
)

func newPost(postBody io.Reader) (Post, error) {
  scanner := bufio.NewScanner(postBody)

  readMetaLine := func(tagName string) string {
    scanner.Scan()
    return strings.TrimPrefix(scanner.Text(), tagName)
}

  title := readMetaLine(titleSeparator)
  description := readMetaLine(descriptionSeparator)
  tags := strings.Split(readMetaLine(tagsSeparator), ", ")

  scanner.Scan() // ignore a line

  buf := bytes.Buffer{}
  for scanner.Scan() {
    fmt.Fprintln(&buf, scanner.Text())
  }
  body := strings.TrimSuffix(buf.String(), "\n")

  return Post{
    Title:       title,
    Description: description,
    Tags:        tags,
    Body:        body,
  }, nil
}

