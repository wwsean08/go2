CREATE SCHEMA "go";
ALTER SCHEMA "go" OWNER TO postgres;

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';

SET SCHEMA 'go';
SET search_path = go, pg_catalog, public;
SET default_tablespace = '';
SET default_with_oids = false;

CREATE TABLE keyword (
  id integer NOT NULL,
  keyword text NOT NULL,
  result_mode integer DEFAULT 0,
  is_regex boolean DEFAULT false NOT NULL
);

ALTER TABLE keyword OWNER TO postgres;

CREATE SEQUENCE keyword_id_seq
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

ALTER TABLE keyword_id_seq OWNER TO postgres;
ALTER SEQUENCE keyword_id_seq OWNED BY keyword.id;

CREATE TABLE keyword_url (
  keyword_id integer NOT NULL,
  url_id integer NOT NULL
);

ALTER TABLE keyword_url OWNER TO postgres;
COMMENT ON TABLE keyword_url IS 'Junction table for keywords to urls and vice versa';

CREATE TABLE url (
  id integer NOT NULL,
  url text NOT NULL,
  click_count integer DEFAULT 0,
  last_clicked timestamp without time zone,
  title text
);

ALTER TABLE url OWNER TO postgres;

CREATE TABLE url_history (
  id integer NOT NULL,
  username character varying(100),
  url_edited integer,
  change_time timestamp with time zone DEFAULT now(),
  change_type text
);

ALTER TABLE url_history OWNER TO postgres;

CREATE SEQUENCE url_history_id_seq
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

ALTER TABLE url_history_id_seq OWNER TO postgres;
ALTER SEQUENCE url_history_id_seq OWNED BY url_history.id;

CREATE SEQUENCE url_id_seq
START WITH 1000
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

ALTER TABLE url_id_seq OWNER TO postgres;
ALTER SEQUENCE url_id_seq OWNED BY url.id;

CREATE TABLE variable (
  id integer NOT NULL,
  _variable text,
  default_value text,
  enabled boolean DEFAULT false
);

ALTER TABLE variable OWNER TO postgres;

CREATE TABLE variable_history (
  id integer NOT NULL,
  username character varying(100),
  variable_edited integer,
  change_time timestamp with time zone DEFAULT now()
);

ALTER TABLE variable_history OWNER TO postgres;

CREATE SEQUENCE variable_history_id_seq
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

ALTER TABLE variable_history_id_seq OWNER TO postgres;
ALTER SEQUENCE variable_history_id_seq OWNED BY variable_history.id;

CREATE SEQUENCE variable_id_seq
START WITH 1
INCREMENT BY 1
NO MINVALUE
NO MAXVALUE
CACHE 1;

ALTER TABLE variable_id_seq OWNER TO postgres;
ALTER SEQUENCE variable_id_seq OWNED BY variable.id;
ALTER TABLE ONLY keyword ALTER COLUMN id SET DEFAULT nextval('keyword_id_seq'::regclass);
ALTER TABLE ONLY url_history ALTER COLUMN id SET DEFAULT nextval('url_history_id_seq'::regclass);
ALTER TABLE ONLY variable ALTER COLUMN id SET DEFAULT nextval('variable_id_seq'::regclass);
ALTER TABLE ONLY variable_history ALTER COLUMN id SET DEFAULT nextval('variable_history_id_seq'::regclass);

ALTER TABLE ONLY url_history
  ADD CONSTRAINT pk_history PRIMARY KEY (id);
ALTER TABLE ONLY variable_history
  ADD CONSTRAINT pk_history_0 PRIMARY KEY (id);
ALTER TABLE ONLY keyword
  ADD CONSTRAINT pk_keyword PRIMARY KEY (id);
ALTER TABLE ONLY keyword_url
  ADD CONSTRAINT pk_keywordurl PRIMARY KEY (keyword_id, url_id);
ALTER TABLE ONLY url
  ADD CONSTRAINT pk_url PRIMARY KEY (id);
ALTER TABLE ONLY variable
  ADD CONSTRAINT pk_variable PRIMARY KEY (id);

CREATE INDEX idx_history ON url_history USING btree (url_edited);
CREATE INDEX idx_history_0 ON variable_history USING btree (variable_edited);
CREATE UNIQUE INDEX keyword_keyword_uindex ON keyword USING btree (keyword);
CREATE UNIQUE INDEX url_url_uindex ON url USING btree (url);

ALTER TABLE ONLY url_history
  ADD CONSTRAINT fk_history FOREIGN KEY (url_edited) REFERENCES url(id);
ALTER TABLE ONLY variable_history
  ADD CONSTRAINT fk_variable_history FOREIGN KEY (variable_edited) REFERENCES variable(id);
ALTER TABLE ONLY keyword_url
  ADD CONSTRAINT keyword_url_keyword_id_fkey FOREIGN KEY (keyword_id) REFERENCES keyword(id);
ALTER TABLE ONLY keyword_url
  ADD CONSTRAINT keyword_url_url_id_fkey FOREIGN KEY (url_id) REFERENCES url(id);