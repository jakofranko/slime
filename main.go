package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	frames   []string
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	slime := []string{
		`
              ░░░░░░░░░░                                                                                   
          ░░░░        ░░░░░░                                                                               
        ░░                  ░░                                                                             
      ░░                    ░░░░                                                                           
    ░░                      ░░░░░░                                                                         
    ░░                        ░░░░                                                                         
  ░░                ░░    ░░  ░░░░░░                                                                       
  ░░                ██░░  ██    ░░░░                                                                       
  ░░                ██░░  ██    ░░░░                                                                       
  ░░            ░░            ░░░░░░                                                                       
  ░░░░░░                      ░░░░░░                                                                       
    ░░░░░░                  ░░░░░░                                                                         
    ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░                                                                         
        ░░░░░░░░░░░░░░░░░░░░░░        
		`,
		`
		             ░░░░░░░░░░░░            
                 ░░░░          ░░░░░░                                                                                                                                                  
               ░░                    ░░                                                                                                                                                
             ░░                      ░░░░                                                                                                                                              
           ░░                        ░░░░░░                                                                                                                                            
           ░░                  ░░    ░░░░░░                                                                                                                                            
         ░░                    ██    ██░░░░░░                                                                                                                                          
         ░░                    ██    ██  ░░░░                                                                                                                                          
         ░░                ░░  ░░    ░░  ░░░░                                                                                                                                          
         ░░                              ░░░░                                                                                                                                          
         ░░░░░░                        ░░░░░░                                                                                                                                          
           ░░░░░░░░░░░░░░░░░░░░░░░░  ░░░░░░                                                                                                                                            
           ░░░░░░░░                ░░░░░░░░                                                                                                                                            
		`,
		`
            ░░░░░░░░░░            
        ░░            ░░                                                                                                                                                            
      ░░                ░░░░                                                                                                                                                        
    ░░                  ░░░░░░                                                                                                                                                      
  ░░              ██  ░░██░░░░░░                                                                                                                                                    
  ░░              ████  ████░░                                                                                                                                                      
░░                  ██  ▒▒██░░░░░░                                                                                                                                                  
░░                  ██    ▓▓  ░░░░                                                                                                                                                  
░░                            ░░░░                                                                                                                                                  
░░                            ░░░░                                                                                                                                                  
░░░░                        ░░░░░░                                                                                                                                                  
  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░                                                                                                                                                    
  ░░░░░░░░              ░░░░░░░░                                                                                                                                                    
  `,
		`
              ░░░░░░░░░░░░                        
          ░░░░            ░░░░                    
                            ░░░░                  
        ░░                    ░░░░                
      ░░                      ░░░░░░              
      ░░                        ░░░░              
      ░░                        ░░░░              
      ░░                        ░░░░              
      ░░                        ░░░░              
    ░░                        ░░░░░░░░            
    ░░                ▓▓    ░░██  ░░░░            
    ░░              ██▓▓    ████  ░░░░            
    ░░              ██░░    ██░░  ░░░░            
    ░░░░░░      ░░              ░░░░░░            
      ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░              
`,
	}

	return model{
		frames:   slime,
		choices:  []string{"one", "two", "three", "four", "five"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	l := m.frames

	// Switch on the type of message received
	switch msg := msg.(type) {
	// Key presses
	case tea.KeyPressMsg:

		// What key was pressed?
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(l)-1 {
				m.cursor++
			}
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to Bubble Team for proccessing
	return m, nil
}

func (m model) View() tea.View {
	s := "Slimey!\n\n"

	for i, choice := range m.frames {

		// No cursor
		cursor := " "
		if m.cursor == i {
			cursor = ">~~"
		}

		// Is the choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "!!!"
		}

		// Render row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	// Off to the UI for rendering
	return tea.NewView(s)
}

func main() {
	// Run spinner in a separate goroutine
	// done := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		default:
	// 			for _, frame := range slime {
	// 				// \r moves cursor to the start of the line
	// 				fmt.Printf("\r%s", frame)
	// 				time.Sleep(100 * time.Millisecond)
	// 			}
	// 		}
	// 	}
	// }()

	// // Simulate work
	// time.Sleep(3 * time.Second)
	// done <- true
	// fmt.Println("\rDone!          ")

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
