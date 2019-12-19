package texttable

import (
	"testing"
)

func TestTextTableFormatter_Output(t *testing.T) {
	tests := []struct {
		name       string
		input      [][]string
		outputFile string
		want       string
	}{
		{
			name: "two rows",
			input: [][]string{
				{
					"Lorem1 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Pharetra pharetra massa massa ultricies mi quis hendrerit dolor magna. Arcu bibendum at varius vel pharetra. Dis parturient montes nascetur ridiculus mus mauris. Id volutpat lacus laoreet non.",
					"Lorem2 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet facilisis magna etiam tempor. Leo in vitae turpis massa sed elementum tempus. Purus semper eget duis at tellus at. Pellentesque habitant morbi tristique senectus et netus et.",
					"Lorem3 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Orci ac auctor augue mauris augue neque. Pellentesque id nibh tortor id aliquet lectus proin nibh nisl. Sociis natoque penatibus et magnis dis parturient montes. Sagittis aliquam malesuada bibendum arcu vitae elementum curabitur vitae.",
				},
				{
					"Lorem4 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Malesuada proin libero nunc consequat interdum varius sit amet mattis. Velit euismod in pellentesque massa placerat duis. Lacinia quis vel eros donec ac odio tempor. Mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing.",
					"Lorem5 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Commodo elit at imperdiet dui accumsan sit amet. Placerat in egestas erat imperdiet sed. Dictum varius duis at consectetur lorem donec massa. Orci phasellus egestas tellus rutrum tellus pellentesque eu.",
					"Lorem6 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus sed vulputate odio ut enim blandit volutpat. Leo vel fringilla est ullamcorper. Phasellus vestibulum lorem sed risus. Congue mauris rhoncus aenean vel elit scelerisque.",
				},
			},
			want: `  Lorem1 ipsum dolor sit amet,      Lorem2 ipsum dolor sit amet,      Lorem3 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor             sed do eiusmod tempor           
  incididunt ut labore et dolore    incididunt ut labore et dolore    incididunt ut labore et dolore  
  magna aliqua. Pharetra            magna aliqua. Amet facilisis      magna aliqua. Orci ac auctor    
  pharetra massa massa ultricies    magna etiam tempor. Leo in        augue mauris augue neque.       
  mi quis hendrerit dolor magna.    vitae turpis massa sed            Pellentesque id nibh tortor id  
  Arcu bibendum at varius vel       elementum tempus. Purus semper    aliquet lectus proin nibh       
  pharetra. Dis parturient          eget duis at tellus at.           nisl. Sociis natoque penatibus  
  montes nascetur ridiculus mus     Pellentesque habitant morbi       et magnis dis parturient        
  mauris. Id volutpat lacus         tristique senectus et netus       montes. Sagittis aliquam        
  laoreet non.                      et.                               malesuada bibendum arcu vitae   
                                                                      elementum curabitur vitae.      
                                                                                                      
  Lorem4 ipsum dolor sit amet,      Lorem5 ipsum dolor sit amet,      Lorem6 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor             sed do eiusmod tempor           
  incididunt ut labore et dolore    incididunt ut labore et dolore    incididunt ut labore et dolore  
  magna aliqua. Malesuada proin     magna aliqua. Commodo elit at     magna aliqua. Risus sed         
  libero nunc consequat interdum    imperdiet dui accumsan sit        vulputate odio ut enim blandit  
  varius sit amet mattis. Velit     amet. Placerat in egestas erat    volutpat. Leo vel fringilla     
  euismod in pellentesque massa     imperdiet sed. Dictum varius      est ullamcorper. Phasellus      
  placerat duis. Lacinia quis       duis at consectetur lorem         vestibulum lorem sed risus.     
  vel eros donec ac odio tempor.    donec massa. Orci phasellus       Congue mauris rhoncus aenean    
  Mauris cursus mattis molestie     egestas tellus rutrum tellus      vel elit scelerisque.           
  a iaculis at erat pellentesque    pellentesque eu.                                                  
  adipiscing.                                                                                         
                                                                                                      `,
			outputFile: "./testdata/two-by-three.txt",
		},
		{
			name: "non constant row length",
			input: [][]string{
				{
					"Lorem1 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Pharetra pharetra massa massa ultricies mi quis hendrerit dolor magna. Arcu bibendum at varius vel pharetra. Dis parturient montes nascetur ridiculus mus mauris. Id volutpat lacus laoreet non.",
					"Lorem2 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet facilisis magna etiam tempor. Leo in vitae turpis massa sed elementum tempus. Purus semper eget duis at tellus at. Pellentesque habitant morbi tristique senectus et netus et.",
				},
				{
					"Lorem4 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Malesuada proin libero nunc consequat interdum varius sit amet mattis. Velit euismod in pellentesque massa placerat duis. Lacinia quis vel eros donec ac odio tempor. Mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing.",
					"Lorem5 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Commodo elit at imperdiet dui accumsan sit amet. Placerat in egestas erat imperdiet sed. Dictum varius duis at consectetur lorem donec massa. Orci phasellus egestas tellus rutrum tellus pellentesque eu.",
					"Lorem6 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus sed vulputate odio ut enim blandit volutpat. Leo vel fringilla est ullamcorper. Phasellus vestibulum lorem sed risus. Congue mauris rhoncus aenean vel elit scelerisque.",
				},
			},
			want: `  Lorem1 ipsum dolor sit amet,      Lorem2 ipsum dolor sit amet,                                      
  consectetur adipiscing elit,      consectetur adipiscing elit,                                      
  sed do eiusmod tempor             sed do eiusmod tempor                                             
  incididunt ut labore et dolore    incididunt ut labore et dolore                                    
  magna aliqua. Pharetra            magna aliqua. Amet facilisis                                      
  pharetra massa massa ultricies    magna etiam tempor. Leo in                                        
  mi quis hendrerit dolor magna.    vitae turpis massa sed                                            
  Arcu bibendum at varius vel       elementum tempus. Purus semper                                    
  pharetra. Dis parturient          eget duis at tellus at.                                           
  montes nascetur ridiculus mus     Pellentesque habitant morbi                                       
  mauris. Id volutpat lacus         tristique senectus et netus                                       
  laoreet non.                      et.                                                               
                                                                                                      
  Lorem4 ipsum dolor sit amet,      Lorem5 ipsum dolor sit amet,      Lorem6 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor             sed do eiusmod tempor           
  incididunt ut labore et dolore    incididunt ut labore et dolore    incididunt ut labore et dolore  
  magna aliqua. Malesuada proin     magna aliqua. Commodo elit at     magna aliqua. Risus sed         
  libero nunc consequat interdum    imperdiet dui accumsan sit        vulputate odio ut enim blandit  
  varius sit amet mattis. Velit     amet. Placerat in egestas erat    volutpat. Leo vel fringilla     
  euismod in pellentesque massa     imperdiet sed. Dictum varius      est ullamcorper. Phasellus      
  placerat duis. Lacinia quis       duis at consectetur lorem         vestibulum lorem sed risus.     
  vel eros donec ac odio tempor.    donec massa. Orci phasellus       Congue mauris rhoncus aenean    
  Mauris cursus mattis molestie     egestas tellus rutrum tellus      vel elit scelerisque.           
  a iaculis at erat pellentesque    pellentesque eu.                                                  
  adipiscing.                                                                                         
                                                                                                      `,
			outputFile: "./testdata/2-then-3.txt",
		},
		{
			name: "new lines 1",
			input: [][]string{
				{
					"Lorem1 ipsum dolor sit amet, consectetur adipiscing elit, \nsed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Pharetra pharetra massa massa ultricies mi quis hendrerit dolor magna. Arcu bibendum at varius vel pharetra. Dis parturient montes nascetur ridiculus mus mauris. Id volutpat lacus laoreet non.",
					"Lorem2 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet facilisis magna etiam tempor. Leo in vitae turpis massa sed elementum \n\ntempus. Purus semper eget duis at tellus at. Pellentesque habitant morbi tristique senectus et netus et.",
					"Lorem3 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Orci ac auctor augue mauris augue neque. Pellentesque id nibh tortor id aliquet lectus proin nibh nisl. Sociis natoque penatibus et magnis dis parturient montes. Sagittis aliquam malesuada bibendum arcu vitae elementum curabitur vitae.",
				},
				{
					"Lorem4 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Malesuada proin libero nunc consequat interdum varius sit amet mattis. Velit euismod in pellentesque massa placerat duis. Lacinia quis vel eros donec ac odio tempor. Mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing.",
					"Lorem5 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Commodo elit at imperdiet dui accumsan sit amet. \n\n\nPlacerat in egestas erat imperdiet sed. Dictum varius duis at consectetur lorem donec massa. Orci phasellus egestas tellus rutrum tellus pellentesque eu.",
					"Lorem6 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus sed vulputate odio ut enim blandit volutpat. Leo vel fringilla est ullamcorper. Phasellus vestibulum \n\n\n\nlorem sed risus. Congue mauris rhoncus aenean vel elit scelerisque.",
				},
			},
			want: `  Lorem1 ipsum dolor sit amet,      Lorem2 ipsum dolor sit amet,      Lorem3 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor             sed do eiusmod tempor           
  incididunt ut labore et dolore    incididunt ut labore et dolore    incididunt ut labore et dolore  
  magna aliqua. Pharetra            magna aliqua. Amet facilisis      magna aliqua. Orci ac auctor    
  pharetra massa massa ultricies    magna etiam tempor. Leo in        augue mauris augue neque.       
  mi quis hendrerit dolor magna.    vitae turpis massa sed            Pellentesque id nibh tortor id  
  Arcu bibendum at varius vel       elementum                         aliquet lectus proin nibh       
  pharetra. Dis parturient                                            nisl. Sociis natoque penatibus  
  montes nascetur ridiculus mus     tempus. Purus semper eget duis    et magnis dis parturient        
  mauris. Id volutpat lacus         at tellus at. Pellentesque        montes. Sagittis aliquam        
  laoreet non.                      habitant morbi tristique          malesuada bibendum arcu vitae   
                                    senectus et netus et.             elementum curabitur vitae.      
                                                                                                      
  Lorem4 ipsum dolor sit amet,      Lorem5 ipsum dolor sit amet,      Lorem6 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor             sed do eiusmod tempor           
  incididunt ut labore et dolore    incididunt ut labore et dolore    incididunt ut labore et dolore  
  magna aliqua. Malesuada proin     magna aliqua. Commodo elit at     magna aliqua. Risus sed         
  libero nunc consequat interdum    imperdiet dui accumsan sit        vulputate odio ut enim blandit  
  varius sit amet mattis. Velit     amet.                             volutpat. Leo vel fringilla     
  euismod in pellentesque massa                                       est ullamcorper. Phasellus      
  placerat duis. Lacinia quis                                         vestibulum                      
  vel eros donec ac odio tempor.    Placerat in egestas erat                                          
  Mauris cursus mattis molestie     imperdiet sed. Dictum varius                                      
  a iaculis at erat pellentesque    duis at consectetur lorem                                         
  adipiscing.                       donec massa. Orci phasellus       lorem sed risus. Congue mauris  
                                    egestas tellus rutrum tellus      rhoncus aenean vel elit         
                                    pellentesque eu.                  scelerisque.                    
                                                                                                      `,
			outputFile: "./testdata/new-lines-1.txt",
		},
		{
			name: "exceed line width 1",
			input: [][]string{
				{
					"Lorem1 ipsum dolor sit amet, consectetur92873645824y2y47828745682894356826435826458 adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Pharetra pharetra massa massa ultricies mi quis hendrerit dolor magna. Arcu bibendum at varius vel pharetra. Dis parturient montes nascetur ridiculus mus mauris. Id volutpat lacus laoreet non.",
					"Lorem2 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet facilisis magna etiam tempor. Leo in vitae turpis massa sed elementum tempus. Purus semper eget duis at tellus at. Pellentesque habitant morbi tristique senectus et netus et.",
					"Lorem3 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Orci ac auctor augue mauris augue neque. Pellentesque id nibh tortor id aliquet lectus proin nibh nisl. Sociis natoque penatibus et magnis dis parturient montes. Sagittis aliquam malesuada bibendum arcu vitae elementum curabitur vitae.",
				},
				{
					"Lorem4 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Malesuada proin libero nunc consequat interdum varius sit amet mattis. Velit euismod in pellentesque massa placerat duis. Lacinia quis vel eros donec ac odio tempor. Mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing.",
					"Lorem5 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt2345264576827435967298345689263459862938456 ut labore et dolore magna aliqua. Commodo elit at imperdiet dui accumsan sit amet. Placerat in egestas erat imperdiet sed. Dictum varius duis at consectetur lorem donec massa. Orci phasellus egestas tellus rutrum tellus pellentesque eu.",
					"Lorem6 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus sed vulputate odio ut enim blandit volutpat. Leo vel fringilla est ullamcorper. Phasellus vestibulum lorem sed risus. Congue mauris rhoncus aenean vel elit scelerisque.",
				},
			},
			want: `  Lorem1 ipsum dolor sit amet, c    Lorem2 ipsum dolor sit amet,      Lorem3 ipsum dolor sit amet,    
  onsectetur92873645824y2y478287    consectetur adipiscing elit,      consectetur adipiscing elit,    
  45682894356826435826458           sed do eiusmod tempor             sed do eiusmod tempor           
  adipiscing elit, sed do           incididunt ut labore et dolore    incididunt ut labore et dolore  
  eiusmod tempor incididunt ut      magna aliqua. Amet facilisis      magna aliqua. Orci ac auctor    
  labore et dolore magna aliqua.    magna etiam tempor. Leo in        augue mauris augue neque.       
  Pharetra pharetra massa massa     vitae turpis massa sed            Pellentesque id nibh tortor id  
  ultricies mi quis hendrerit       elementum tempus. Purus semper    aliquet lectus proin nibh       
  dolor magna. Arcu bibendum at     eget duis at tellus at.           nisl. Sociis natoque penatibus  
  varius vel pharetra. Dis          Pellentesque habitant morbi       et magnis dis parturient        
  parturient montes nascetur        tristique senectus et netus       montes. Sagittis aliquam        
  ridiculus mus mauris. Id          et.                               malesuada bibendum arcu vitae   
  volutpat lacus laoreet non.                                         elementum curabitur vitae.      
                                                                                                      
  Lorem4 ipsum dolor sit amet,      Lorem5 ipsum dolor sit amet,      Lorem6 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor incididu    sed do eiusmod tempor           
  incididunt ut labore et dolore    nt2345264576827435967298345689    incididunt ut labore et dolore  
  magna aliqua. Malesuada proin     263459862938456 ut labore et      magna aliqua. Risus sed         
  libero nunc consequat interdum    dolore magna aliqua. Commodo      vulputate odio ut enim blandit  
  varius sit amet mattis. Velit     elit at imperdiet dui accumsan    volutpat. Leo vel fringilla     
  euismod in pellentesque massa     sit amet. Placerat in egestas     est ullamcorper. Phasellus      
  placerat duis. Lacinia quis       erat imperdiet sed. Dictum        vestibulum lorem sed risus.     
  vel eros donec ac odio tempor.    varius duis at consectetur        Congue mauris rhoncus aenean    
  Mauris cursus mattis molestie     lorem donec massa. Orci           vel elit scelerisque.           
  a iaculis at erat pellentesque    phasellus egestas tellus                                          
  adipiscing.                       rutrum tellus pellentesque eu.                                    
                                                                                                      `,
			outputFile: "./testdata/exceed-line-width-1.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttf := New(tt.input, Config{
				ColumnWidth:  30,
				ColumnMargin: 2,
				RowMargin:    1,
			})

			output, err := ttf.Output()
			if err != nil {
				t.Fatalf("failed to format text table: %v", err)
			}

			//f, _ := os.Create(tt.outputFile)
			//f.Write([]byte(output))

			if output != tt.want {
				t.Fatalf("expected: \n%v\n\ngot: \n%v\n", tt.want, output)
			}

		})
	}

}

func TestTextTableFormatter_Output_IgnoreNewLines(t *testing.T) {
	tests := []struct {
		name       string
		input      [][]string
		outputFile string
		want       string
	}{
		{
			name: "new lines 1",
			input: [][]string{
				{
					"Lorem1 ipsum dolor sit amet, consectetur adipiscing elit, \nsed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Pharetra pharetra massa massa ultricies mi quis hendrerit dolor magna. Arcu bibendum at varius vel pharetra. Dis parturient montes nascetur ridiculus mus mauris. Id volutpat lacus laoreet non.",
					"Lorem2 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet facilisis magna etiam tempor. Leo in vitae turpis massa sed elementum \n\ntempus. Purus semper eget duis at tellus at. Pellentesque habitant morbi tristique senectus et netus et.",
					"Lorem3 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Orci ac auctor augue mauris augue neque. Pellentesque id nibh tortor id aliquet lectus proin nibh nisl. Sociis natoque penatibus et magnis dis parturient montes. Sagittis aliquam malesuada bibendum arcu vitae elementum curabitur vitae.",
				},
				{
					"Lorem4 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Malesuada proin libero nunc consequat interdum varius sit amet mattis. Velit euismod in pellentesque massa placerat duis. Lacinia quis vel eros donec ac odio tempor. Mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing.",
					"Lorem5 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Commodo elit at imperdiet dui accumsan sit amet. \n\n\nPlacerat in egestas erat imperdiet sed. Dictum varius duis at consectetur lorem donec massa. Orci phasellus egestas tellus rutrum tellus pellentesque eu.",
					"Lorem6 ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus sed vulputate odio ut enim blandit volutpat. Leo vel fringilla est ullamcorper. Phasellus vestibulum \n\n\n\nlorem sed risus. Congue mauris rhoncus aenean vel elit scelerisque.",
				},
			},
			want: `  Lorem1 ipsum dolor sit amet,      Lorem2 ipsum dolor sit amet,      Lorem3 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor             sed do eiusmod tempor           
  incididunt ut labore et dolore    incididunt ut labore et dolore    incididunt ut labore et dolore  
  magna aliqua. Pharetra            magna aliqua. Amet facilisis      magna aliqua. Orci ac auctor    
  pharetra massa massa ultricies    magna etiam tempor. Leo in        augue mauris augue neque.       
  mi quis hendrerit dolor magna.    vitae turpis massa sed            Pellentesque id nibh tortor id  
  Arcu bibendum at varius vel       elementum tempus. Purus semper    aliquet lectus proin nibh       
  pharetra. Dis parturient          eget duis at tellus at.           nisl. Sociis natoque penatibus  
  montes nascetur ridiculus mus     Pellentesque habitant morbi       et magnis dis parturient        
  mauris. Id volutpat lacus         tristique senectus et netus       montes. Sagittis aliquam        
  laoreet non.                      et.                               malesuada bibendum arcu vitae   
                                                                      elementum curabitur vitae.      
                                                                                                      
  Lorem4 ipsum dolor sit amet,      Lorem5 ipsum dolor sit amet,      Lorem6 ipsum dolor sit amet,    
  consectetur adipiscing elit,      consectetur adipiscing elit,      consectetur adipiscing elit,    
  sed do eiusmod tempor             sed do eiusmod tempor             sed do eiusmod tempor           
  incididunt ut labore et dolore    incididunt ut labore et dolore    incididunt ut labore et dolore  
  magna aliqua. Malesuada proin     magna aliqua. Commodo elit at     magna aliqua. Risus sed         
  libero nunc consequat interdum    imperdiet dui accumsan sit        vulputate odio ut enim blandit  
  varius sit amet mattis. Velit     amet. Placerat in egestas erat    volutpat. Leo vel fringilla     
  euismod in pellentesque massa     imperdiet sed. Dictum varius      est ullamcorper. Phasellus      
  placerat duis. Lacinia quis       duis at consectetur lorem         vestibulum lorem sed risus.     
  vel eros donec ac odio tempor.    donec massa. Orci phasellus       Congue mauris rhoncus aenean    
  Mauris cursus mattis molestie     egestas tellus rutrum tellus      vel elit scelerisque.           
  a iaculis at erat pellentesque    pellentesque eu.                                                  
  adipiscing.                                                                                         
                                                                                                      `,
			outputFile: "./testdata/new-lines-1.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttf := New(tt.input, Config{
				ColumnWidth:    30,
				ColumnMargin:   2,
				RowMargin:      1,
				IgnoreNewLines: true,
			})

			output, err := ttf.Output()
			if err != nil {
				t.Fatalf("failed to format text table: %v", err)
			}

			//f, _ := os.Create(tt.outputFile)
			//f.Write([]byte(output))

			if output != tt.want {
				t.Fatalf("expected: \n%v\n\ngot: \n%v\n", tt.want, output)
			}

		})
	}

}
