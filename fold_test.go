package ascii

import "testing"

func TestFold(t *testing.T) {
	tests := []struct{ in, want string }{
		{"æøå", "\x00E6\x00F8\x00E5"}, // 0
		{"öl", "ol"},                  // 1
		{"Çiğdem Çiçek Şarkı Sözü", "Cigdem Cicek Sarki Sozu"},   // 2
		{"Des mot clés À LA CHAÎNE", "Des mot cles A LA CHAINE"}, // 3
		{"Bāzār'hā-yi Īrān", "Bazar'ha-yi Iran"},                 // 4
	}

	for i, test := range tests {
		if got := Fold(test.in); got != test.want {
			t.Errorf("%d: Fold(%q) == %q; want %q", i, test.in, got, test.want)
		}
	}
}
