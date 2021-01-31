package quotes

import (
	"fmt"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

// DB is a quote database.
type DB struct {
	db *bolt.DB
}

const (
	quoteBucket = "standard"
    bucketName = "quoteBucket"
)

// Open opens the database file at path and returns a DB or an error.
func Open(path string) (*DB, error) {

	// TODO:
	// Open the DB file at path using bolt.Open().
	// Pass 0600 as file mode, and nil as options.
	// Return a pointer to the open DB, or an error.

	open, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db")
	}

	db := DB{open}
	return &db, err

}

func (d *DB) Close() error {

	// TODO:
	// Close the database d.db and return any
	// error or nil.

	err := d.db.Close()
	return err
}

// Create takes a quote and saves it to the database, using the author name
// as the key. If the author already exists, Create returns an error.
func (d *DB) Create(q *Quote) error {
	err := d.db.Update(func(tx *bolt.Tx) error {

		// TODO: Create a bucket if it does not exist already.
		// Use the constant quoteBucket as the bucket name.
		//
		// Remember to use []byte(...) to convert a string into a byte
		// slice if required.


		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return errors.Wrap(err, "failed to create bolt bucket")
		}

		// Ensure that the quote we want to save does not already exist.
		// Hint: Call bucket.Get and verify if the result has zero length.

		if value := bucket.Get([]byte(q.Author)); value != nil {
			return fmt.Errorf("giving up creating quote for author '%s' as already exists", q.Author)
		}

		// Serialize the quote, using the Serialize method from quote.go.
		qbits, err := q.Serialize()

		// Put the serialized quote into the bucket.
		err = bucket.Put([]byte(q.Author), qbits)
		if err != nil {
			return errors.Wrap(err, "failed writing quote to bolt")
		}

		return nil
	})

	// TODO: Check the error returned by d.db.Update. Return an error or nil.
	return err
}

// Get takes an author name and retrieves the corresponding quote from the DB.
func (d *DB) Get(author string) (*Quote, error) {
	q := &Quote{}
	err := d.db.View(func(tx *bolt.Tx) error {

		// Get the bucket with the name as specified by the constant quoteBucket.
		// The bucket is available from the transaction object tx.
		bucket := tx.Bucket([]byte(bucketName))

		// Get the desired quote by the author's name.
		//
		// Again, remember to use []byte(...) to convert a string into a byte
		// slice if required.
		qbits := bucket.Get([]byte(author))

		// Deserialize the quote into q - use the Deserialize method from
		// quote.go for this.
		// Remember that we are within a closure that has access to q.
		err := q.Deserialize(qbits)

		// Check and return any error that occurs.
		return err
	})
	// Check the error returned by d.db.View.
	// Return (nil, err) or (&q, nil), respectively.
	return q, err
}

// List lists all records in the DB.
func (d *DB) List() ([]*Quote, error) {
	// The database returns byte slices that we need to de-serialize
	// into Quote structures.
	structList := []*Quote{}

	// We use a View as we don't update anything.
	err := d.db.View(func(tx *bolt.Tx) error {

		// Get the bucket from the transaction tx.
		bucket := tx.Bucket([]byte(bucketName))

		// Iterate over all elements of the bucket.
		// Hint: BoltDB has a ForEach method for this.
		//   * For each element, create a new *Quote and deserialize
		//     the element value into the *Quote.
		//   * Then append the *Quote to structList.
		err := bucket.ForEach(func(key []byte, value []byte) error {
			q := &Quote{}
			err := q.Deserialize(value)
			structList = append(structList, q)
			return err
		})

		return err
	})

	// TODO: Check the error returned by d.db.View().
	// Return (structList, nil) or (nil, err), respectively.
	if err == nil {
		return structList, nil
	} else {
		return nil, err
	}
}
