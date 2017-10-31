package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

func AddGigasecond( t time.Time) time.Time {
	gigasecond := 1000000000
        return  t.Add( time.Duration(gigasecond) * time.Second )
}
