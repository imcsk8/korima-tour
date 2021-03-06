--
-- PostgreSQL database dump
--

-- Dumped from database version 10.6
-- Dumped by pg_dump version 10.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: acl; Type: TABLE; Schema: public; Owner: korima_pg
--

CREATE TABLE public.acl (
    id uuid NOT NULL,
    level integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.acl OWNER TO korima_pg;

--
-- Name: band_venue_event_requests; Type: TABLE; Schema: public; Owner: korima_pg
--

CREATE TABLE public.band_venue_event_requests (
    id uuid NOT NULL,
    band_id uuid NOT NULL,
    venue_id uuid NOT NULL,
    request_date date NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    owner_id uuid NOT NULL,
    status integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.band_venue_event_requests OWNER TO korima_pg;

--
-- Name: bands; Type: TABLE; Schema: public; Owner: korima_pg
--

CREATE TABLE public.bands (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    owner_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    photo character varying(255) DEFAULT ''::character varying,
    country integer DEFAULT 1,
    state integer DEFAULT 9,
    city integer DEFAULT 15,
    location_text character varying(255) DEFAULT ''::character varying
);


ALTER TABLE public.bands OWNER TO korima_pg;

--
-- Name: bookers; Type: TABLE; Schema: public; Owner: korima_pg
--

CREATE TABLE public.bookers (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    owner_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    photo character varying(255) DEFAULT ''::character varying,
    country integer DEFAULT 1,
    state integer DEFAULT 9,
    city integer DEFAULT 15,
    location_text character varying(255) DEFAULT ''::character varying
);


ALTER TABLE public.bookers OWNER TO korima_pg;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: korima_pg
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO korima_pg;

--
-- Name: users; Type: TABLE; Schema: public; Owner: korima_pg
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    firstname character varying(255),
    middlename character varying(255),
    lastname character varying(255),
    mlastname character varying(255),
    email character varying(255) NOT NULL,
    phone character varying(255),
    admin boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    password_hash character varying(255) DEFAULT ''::character varying NOT NULL,
    username character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.users OWNER TO korima_pg;

--
-- Name: venues; Type: TABLE; Schema: public; Owner: korima_pg
--

CREATE TABLE public.venues (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    owner_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    photo character varying(255) DEFAULT ''::character varying,
    country integer,
    state integer,
    city integer,
    location_text character varying(255) DEFAULT ''::character varying NOT NULL,
    address character varying(255) DEFAULT ''::character varying NOT NULL,
    zip character varying(255) DEFAULT ''::character varying NOT NULL,
    phone character varying(255) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.venues OWNER TO korima_pg;

--
-- Name: acl acl_pkey; Type: CONSTRAINT; Schema: public; Owner: korima_pg
--

ALTER TABLE ONLY public.acl
    ADD CONSTRAINT acl_pkey PRIMARY KEY (id);


--
-- Name: band_venue_event_requests band_venue_event_request_pkey; Type: CONSTRAINT; Schema: public; Owner: korima_pg
--

ALTER TABLE ONLY public.band_venue_event_requests
    ADD CONSTRAINT band_venue_event_request_pkey PRIMARY KEY (id);


--
-- Name: bands bands_pkey; Type: CONSTRAINT; Schema: public; Owner: korima_pg
--

ALTER TABLE ONLY public.bands
    ADD CONSTRAINT bands_pkey PRIMARY KEY (id);


--
-- Name: bookers bookers_pkey; Type: CONSTRAINT; Schema: public; Owner: korima_pg
--

ALTER TABLE ONLY public.bookers
    ADD CONSTRAINT bookers_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: korima_pg
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: venues venues_pkey; Type: CONSTRAINT; Schema: public; Owner: korima_pg
--

ALTER TABLE ONLY public.venues
    ADD CONSTRAINT venues_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: korima_pg
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

