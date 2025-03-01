// internal/cli/formatter/table.go
package formatter

import (
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/layla-lili/blockchain_tools/pkg/types"
)

// TableFormatter formats data as a table
type TableFormatter struct{}

// NewTableFormatter creates a new table formatter
func NewTableFormatter() *TableFormatter {
	return &TableFormatter{}
}

// Format formats the data as a table
func (f *TableFormatter) Format(w io.Writer, data interface{}) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	// Handle different data types
	switch v := data.(type) {
	case *types.Block:
		return f.formatBlock(tw, v)
	case *types.Transaction:
		return f.formatTransaction(tw, v)
	case []*types.Block:
		return f.formatBlocks(tw, v)
	case []*types.Transaction:
		return f.formatTransactions(tw, v)
	default:
		return fmt.Errorf("unsupported data type: %T", data)
	}
}

// formatBlock formats a single block
func (f *TableFormatter) formatBlock(w io.Writer, block *types.Block) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	fmt.Fprintln(tw, "BLOCK INFORMATION:")
	fmt.Fprintf(tw, "Hash:\t%s\n", block.Hash)
	fmt.Fprintf(tw, "Height:\t%d\n", block.Height)
	fmt.Fprintf(tw, "Previous Hash:\t%s\n", block.PreviousHash)
	fmt.Fprintf(tw, "Timestamp:\t%d\n", block.Timestamp)
	fmt.Fprintf(tw, "Transactions:\t%d\n", len(block.Transactions))
	fmt.Fprintf(tw, "Size:\t%d bytes\n", block.Size)

	if len(block.Transactions) > 0 {
		fmt.Fprintln(tw, "\nTRANSACTIONS:")
		fmt.Fprintln(tw, "HASH\tFROM\tTO\tVALUE")

		for _, tx := range block.Transactions {
			fmt.Fprintf(tw, "%s\t%s\t%s\t%d\n",
				tx.Hash,
				truncateString(tx.From, 8),
				truncateString(tx.To, 8),
				tx.Value)
		}
	}

	return tw.Flush()
}

// formatTransaction formats a single transaction
func (f *TableFormatter) formatTransaction(w io.Writer, tx *types.Transaction) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	fmt.Fprintln(tw, "TRANSACTION INFORMATION:")
	fmt.Fprintf(tw, "Hash:\t%s\n", tx.Hash)
	fmt.Fprintf(tw, "From:\t%s\n", tx.From)
	fmt.Fprintf(tw, "To:\t%s\n", tx.To)
	fmt.Fprintf(tw, "Value:\t%d\n", tx.Value)
	fmt.Fprintf(tw, "Status:\t%s\n", tx.Status)
	fmt.Fprintf(tw, "Block Hash:\t%s\n", tx.BlockHash)
	fmt.Fprintf(tw, "Timestamp:\t%d\n", tx.Timestamp)

	if len(tx.Data) > 0 {
		fmt.Fprintf(tw, "Data:\t%x\n", tx.Data)
	}

	return tw.Flush()
}

// formatTransactions formats multiple transactions
func (f *TableFormatter) formatTransactions(w io.Writer, txs []*types.Transaction) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	fmt.Fprintln(tw, "HASH\tFROM\tTO\tVALUE\tSTATUS")

	for _, tx := range txs {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%d\t%s\n",
			truncateString(tx.Hash, 12),
			truncateString(tx.From, 8),
			truncateString(tx.To, 8),
			tx.Value,
			tx.Status)
	}

	return tw.Flush()
}

// formatBlocks formats multiple blocks in a table format
func (f *TableFormatter) formatBlocks(w io.Writer, blocks []*types.Block) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	// Print header
	fmt.Fprintln(tw, "BLOCKS:")
	fmt.Fprintln(tw, "HEIGHT\tHASH\tTIMESTAMP\tTX COUNT\tSIZE")

	// Print each block as a row
	for _, block := range blocks {
		fmt.Fprintf(tw, "%d\t%s\t%s\t%d\t%d bytes\n",
			block.Height,
			truncateString(block.Hash, 12),
			time.Unix(block.Timestamp, 0).Format(time.RFC3339),
			len(block.Transactions),
			block.Size)
	}

	return tw.Flush()
}

// Helper function to truncate long strings (like hashes and addresses)
func truncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

// formatBlockCount formats the block count information
func (f *TableFormatter) formatBlockCount(w io.Writer, data interface{}) error {
	// Create a new tabwriter that wraps the io.Writer
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	if resp, ok := data.(struct {
		BlockHeight uint64 `json:"blockHeight"`
		Timestamp   int64  `json:"timestamp"`
	}); ok {
		fmt.Fprintln(tw, "BLOCKCHAIN STATUS:")
		fmt.Fprintf(tw, "Current Height:\t%d\n", resp.BlockHeight)
		fmt.Fprintf(tw, "Last Updated:\t%s\n", time.Unix(resp.Timestamp, 0))
		return tw.Flush()
	}
	return fmt.Errorf("invalid data type for block count")
}
