package cli

import (
	"testing"
)



func TestApp_Run(t *testing.T){
	tests:=[]struct{
		name string
		args []string
		wantErr bool
	}{
		{
			name:    "no arguments",
			args:    []string{"todo"},
			wantErr: true,
		},
		{
			name:    "invalid command",
			args:    []string{"todo", "invalid"},
			wantErr: true,
		},
		{
			name:    "list command",
			args:    []string{"todo", "list"},
			wantErr: false,
		},
		{
			name:    "add command without description",
			args:    []string{"todo", "add"},
			wantErr: true,
		},
		{
			name:    "add command with description",
			args:    []string{"todo", "add", "Test task"},
			wantErr: false,
		},
		//help and version commands
		{
			name:    "help command",
			args:    []string{"todo", "-h"},
			wantErr: false,
		},
		{
			name:    "version command",
			args:    []string{"todo", "-v"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewApp("", "", "")
			err := a.Run(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}