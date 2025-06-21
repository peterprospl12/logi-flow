-- users
CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       role TEXT NOT NULL
);

-- orders
CREATE TABLE orders (
                        id UUID PRIMARY KEY,
                        shipper_id UUID REFERENCES users(id),
                        "from" TEXT NOT NULL,
                        "to" TEXT NOT NULL,
                        price NUMERIC NOT NULL,
                        status TEXT NOT NULL
);

-- bids
CREATE TABLE bids (
                      id UUID PRIMARY KEY,
                      order_id UUID REFERENCES orders(id),
                      carrier_id UUID REFERENCES users(id),
                      price NUMERIC NOT NULL,
                      status TEXT NOT NULL
);

-- order_events
CREATE TABLE order_events (
                              id UUID PRIMARY KEY,
                              order_id UUID REFERENCES orders(id),
                              event_type TEXT NOT NULL,
                              created_at TIMESTAMP NOT NULL DEFAULT now(),
                              payload JSONB
);