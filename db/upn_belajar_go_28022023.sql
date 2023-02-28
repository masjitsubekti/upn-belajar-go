--
-- PostgreSQL database dump
--

-- Dumped from database version 10.21 (Ubuntu 10.21-1.pgdg22.04+1)
-- Dumped by pg_dump version 10.21 (Ubuntu 10.21-1.pgdg22.04+1)

-- Started on 2023-02-28 11:31:39 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 13104)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2996 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 17022)
-- Name: jenis_mitra; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.jenis_mitra (
    id character varying(36) NOT NULL,
    nama_jenis_mitra character varying(200),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.jenis_mitra OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 17034)
-- Name: m_kelas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_kelas (
    id character varying(36) NOT NULL,
    kode character varying(100),
    nama character varying(100),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.m_kelas OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 17028)
-- Name: m_siswa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_siswa (
    id character varying(36) NOT NULL,
    nama character varying(100),
    kelas character varying(100),
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    created_by character varying(36),
    updated_by character varying(36),
    is_deleted boolean DEFAULT false
);


ALTER TABLE public.m_siswa OWNER TO postgres;

--
-- TOC entry 2986 (class 0 OID 17022)
-- Dependencies: 196
-- Data for Name: jenis_mitra; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2988 (class 0 OID 17034)
-- Dependencies: 198
-- Data for Name: m_kelas; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.m_kelas (id, kode, nama, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('54c9cd72-9541-4dc4-8864-16dc8236cfcc', 'X-IPA-B', 'X IPA B', '2023-02-28 08:26:53.492433', '2023-02-28 08:26:53.492798', '', NULL, false);
INSERT INTO public.m_kelas (id, kode, nama, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('279db6e8-6f00-4f20-8c22-a5ce914c878b', 'X-IPA-A', 'X IPA A', '2023-02-28 08:25:41.270441', '2023-02-28 08:30:36.543048', '', '', false);


--
-- TOC entry 2987 (class 0 OID 17028)
-- Dependencies: 197
-- Data for Name: m_siswa; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.m_siswa (id, nama, kelas, created_at, updated_at, created_by, updated_by, is_deleted) VALUES ('5428ec3a-4a8e-444b-aa16-dd8f9d2cb0ee', 'Bambang Tri UPDATE', 'A UPDATE', '2023-02-28 11:08:22.69835', '2023-02-28 11:19:48.458771', '', '', true);


-- Completed on 2023-02-28 11:31:39 WIB

--
-- PostgreSQL database dump complete
--

