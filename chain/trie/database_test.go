package trie

import (
	"testing"

	"github.com/skaindev/skaina/neatdb/memorydb"
	"github.com/skaindev/skaina/utilities/common"
)

func TestDatabaseMetarootFetch(t *testing.T) {
	db := NewDatabase(memorydb.New())
	if _, err := db.Node(common.Hash{}); err == nil {
		t.Fatalf("metaroot retrieval succeeded")
	}
}
