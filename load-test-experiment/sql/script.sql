CREATE TABLE IF NOT EXISTS light_table
(
    id          serial PRIMARY KEY,
    field_one   VARCHAR(128) NOT NULL,
    field_two   float        NOT NULL,
    field_three VARCHAR(128) NULL,
    field_four  timestamptz  NULL
);

CREATE TABLE IF NOT EXISTS medium_small_table
(
    id              serial PRIMARY KEY,
    field_one       VARCHAR(128) NOT NULL,
    field_two       float        NOT NULL,
    field_three     VARCHAR(128) NULL,
    field_four      timestamptz  NULL,
    small_large_key int          NOT NULL
);

CREATE TABLE IF NOT EXISTS medium_large_table
(
    id             serial PRIMARY KEY,
    field_one      VARCHAR(128) NOT NULL,
    field_two      float        NOT NULL,
    field_three    VARCHAR(128) NULL,
    field_four     timestamptz  NOT NULL DEFAULT Now(),
    field_five     int          NOT NULL,
    field_six      VARCHAR(128) NOT NULL,
    field_seven    float        NOT NULL,
    field_eight    VARCHAR(128) NULL,
    field_nine     timestamptz  NOT NULL DEFAULT Now(),
    field_ten      int          NOT NULL,
    field_eleven   VARCHAR(128) NOT NULL,
    field_twelve   float        NOT NULL,
    field_thirteen VARCHAR(128) NULL,
    field_fourteen timestamptz  NOT NULL DEFAULT Now()
);

CREATE TABLE IF NOT EXISTS heavy_fourth_table
(
    id          serial PRIMARY KEY,
    field_one   VARCHAR(128) NOT NULL,
    field_two   float        NOT NULL,
    field_three VARCHAR(128) NULL,
    field_four  timestamptz  NULL
);

CREATE TABLE IF NOT EXISTS heavy_third_table
(
    id          serial PRIMARY KEY,
    field_one   VARCHAR(128) NOT NULL,
    field_two   float        NOT NULL,
    field_three VARCHAR(128) NULL,
    field_four  timestamptz  NULL
);

CREATE TABLE IF NOT EXISTS heavy_second_table
(
    id                serial PRIMARY KEY,
    field_one         VARCHAR(128) NOT NULL,
    field_two         float        NOT NULL,
    field_three       VARCHAR(128) NULL,
    field_four        timestamptz  NOT NULL DEFAULT Now(),
    field_five        int          NOT NULL,
    field_six         VARCHAR(128) NOT NULL,
    field_seven       float        NOT NULL,
    field_eight       VARCHAR(128) NULL,
    field_nine        timestamptz  NOT NULL DEFAULT Now(),
    field_ten         int          NOT NULL,
    field_eleven      VARCHAR(128) NOT NULL,
    field_twelve      float        NOT NULL,
    field_thirteen    VARCHAR(128) NULL,
    field_fourteen    timestamptz  NOT NULL DEFAULT Now(),
    second_fourth_key int          NOT NULL,
    second_third_key  int          NOT NULL
);

CREATE TABLE IF NOT EXISTS heavy_first_table
(
    id             serial PRIMARY KEY,
    field_one      VARCHAR(128) NOT NULL,
    field_two      float        NOT NULL,
    field_three    VARCHAR(128) NULL,
    field_four     timestamptz  NOT NULL DEFAULT Now(),
    field_five     int          NOT NULL,
    field_six      VARCHAR(128) NOT NULL,
    field_seven    float        NOT NULL,
    field_eight    VARCHAR(128) NULL,
    field_nine     timestamptz  NOT NULL DEFAULT Now(),
    field_ten      int          NOT NULL,
    field_eleven   VARCHAR(128) NOT NULL,
    field_twelve   float        NOT NULL,
    field_thirteen VARCHAR(128) NULL,
    field_fourteen timestamptz  NOT NULL DEFAULT Now(),
    first_second   int          NOT NULL
);