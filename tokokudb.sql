PGDMP  (    +                |            tokokudb    16.4    16.4                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    82719    tokokudb    DATABASE     �   CREATE DATABASE tokokudb WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE tokokudb;
                postgres    false            �            1259    82815    carts    TABLE       CREATE TABLE public.carts (
    user_id integer NOT NULL,
    product_id integer NOT NULL,
    id integer NOT NULL,
    quantity integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.carts;
       public         heap    postgres    false            �            1259    82821    carts_id_seq    SEQUENCE     �   ALTER TABLE public.carts ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.carts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    221            �            1259    82791    products    TABLE     E  CREATE TABLE public.products (
    store_id integer NOT NULL,
    id integer NOT NULL,
    product_name character varying(255),
    price numeric,
    description text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    url_image text
);
    DROP TABLE public.products;
       public         heap    postgres    false            �            1259    82814    products_id_seq    SEQUENCE     �   ALTER TABLE public.products ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    219            �            1259    82777    stores    TABLE     �   CREATE TABLE public.stores (
    user_id integer NOT NULL,
    id integer NOT NULL,
    name character varying(255),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.stores;
       public         heap    postgres    false            �            1259    82787    stores_id_seq    SEQUENCE     �   ALTER TABLE public.stores ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.stores_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    217            �            1259    82720    users    TABLE     �  CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255),
    name character varying(255),
    phone character varying(255),
    address character varying(255),
    store_status boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    82723    users_id_seq    SEQUENCE     �   ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public          postgres    false    215                      0    82815    carts 
   TABLE DATA           Z   COPY public.carts (user_id, product_id, id, quantity, created_at, updated_at) FROM stdin;
    public          postgres    false    221   �!       	          0    82791    products 
   TABLE DATA           u   COPY public.products (store_id, id, product_name, price, description, created_at, updated_at, url_image) FROM stdin;
    public          postgres    false    219   �"                 0    82777    stores 
   TABLE DATA           K   COPY public.stores (user_id, id, name, created_at, updated_at) FROM stdin;
    public          postgres    false    217   f$                 0    82720    users 
   TABLE DATA           p   COPY public.users (id, email, password, name, phone, address, store_status, created_at, updated_at) FROM stdin;
    public          postgres    false    215   �$                  0    0    carts_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.carts_id_seq', 21, true);
          public          postgres    false    222                       0    0    products_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.products_id_seq', 8, true);
          public          postgres    false    220                       0    0    stores_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.stores_id_seq', 5, true);
          public          postgres    false    218                       0    0    users_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.users_id_seq', 20, true);
          public          postgres    false    216            i           2606    82736    users email 
   CONSTRAINT     G   ALTER TABLE ONLY public.users
    ADD CONSTRAINT email UNIQUE (email);
 5   ALTER TABLE ONLY public.users DROP CONSTRAINT email;
       public            postgres    false    215            q           2606    82819    carts id 
   CONSTRAINT     F   ALTER TABLE ONLY public.carts
    ADD CONSTRAINT id PRIMARY KEY (id);
 2   ALTER TABLE ONLY public.carts DROP CONSTRAINT id;
       public            postgres    false    221            o           2606    82797    products products_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public            postgres    false    219            m           2606    82786    stores storeid 
   CONSTRAINT     L   ALTER TABLE ONLY public.stores
    ADD CONSTRAINT storeid PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.stores DROP CONSTRAINT storeid;
       public            postgres    false    217            k           2606    82730    users userid 
   CONSTRAINT     J   ALTER TABLE ONLY public.users
    ADD CONSTRAINT userid PRIMARY KEY (id);
 6   ALTER TABLE ONLY public.users DROP CONSTRAINT userid;
       public            postgres    false    215            t           2606    82827    carts product_id    FK CONSTRAINT        ALTER TABLE ONLY public.carts
    ADD CONSTRAINT product_id FOREIGN KEY (product_id) REFERENCES public.products(id) NOT VALID;
 :   ALTER TABLE ONLY public.carts DROP CONSTRAINT product_id;
       public          postgres    false    219    4719    221            s           2606    82809    products store_id    FK CONSTRAINT     |   ALTER TABLE ONLY public.products
    ADD CONSTRAINT store_id FOREIGN KEY (store_id) REFERENCES public.stores(id) NOT VALID;
 ;   ALTER TABLE ONLY public.products DROP CONSTRAINT store_id;
       public          postgres    false    4717    219    217            r           2606    82780    stores user_id    FK CONSTRAINT     w   ALTER TABLE ONLY public.stores
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.users(id) NOT VALID;
 8   ALTER TABLE ONLY public.stores DROP CONSTRAINT user_id;
       public          postgres    false    215    217    4715            u           2606    82822    carts user_id    FK CONSTRAINT     v   ALTER TABLE ONLY public.carts
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.users(id) NOT VALID;
 7   ALTER TABLE ONLY public.carts DROP CONSTRAINT user_id;
       public          postgres    false    4715    221    215               �   x�}�;!��z=E��2� �Yr�s�M,����O|P=� :�`d9�O����&�����(��P�ڑ��ԩԬH������e�S����
��z�(�:�)��jU��o�J�ʵ�Tmw�{��%E�����YY�TzVڸB k��Q�u*�(8�ҀZs�UI��0fE;0	q�%թ�w��@���K�Z^PJy�g�G      	   �  x����j1����}Yit\i����B�B��=؊���R���]����@
��G��͈!�����!k���f������mqBs_gܴ�|�>�p�p�ā�l�J�Le�Q��կ�|B]�SEiʤicH>�u�!�:�����	n��R=)�=�J]��M1e����w�.�
�#�����l<HQ)�l^3aJ�ڣf� ZH�ً�ʇmr��w�TRV��eB����"��(S	�ua�s�#8������&�6�	�Eu�k�S����n�>�2胫�n�����x��o����v��7i�����>BU�����_���¥�f�	}t���C��~�W���������G�[�ٵ��=����b ��R�R�^�����M�ւ�� �����B�Nt|"�I��_Qe0�         a   x�34�4����WpLJL��4202�5��54S04�21�24�33��42�60�/�eh�i
1˻4/3/��R��������R�������0l�\1z\\\ �� �         k  x���Ks�8����`��L]=lɫ�$�c܄@B�F��?x�׏�'=	!3�)�����;�G����c�}�R���VV�@���a�T�f��Js����}��@�E��1�@�?���tuS˃,��j��=o���MWd�!0����:�9G!$3�*�A<x�����n��cӉ;c�������!-�,"�~��תL���d���T7�o����;�̅�����O7=��F�bq��tK��k�a�lɑٷ�΁�]��ݭ�-����X�kOyy�U����m���H���^����0�Es��!ת:�7! ץ<��,P
ͽp�w������|���u[���2k��
=���	<�C�)�`qna1?
�\��8G�*�9�wls�8�ٽ���y(zi&��|N���b�Kw�ق�u��~��	R�_�M�A,�,�Q���y	�����ej�l��/cn:�2z/�L�k��]_Ʋ׵��a�󨑈�J�ɷ'��������~ը�
�W����!��� ~�� 45]'4�>"j���f�c<���̐},&q8�2��DuG�� {\�M�	���]�U�J*��l�@��H�3�EU��k?�s�n��Σs[�N�Ҙ5�bQ��>����b��C<��6�[���l���J��
�5���i��j}uT��JU��� �0�
(210F.���d���	z�'��S���Y�j$)�l��b[�}������/�5�r�z�f8.�7="�R��e|a��0X�J�!�K�OU��_%�	G��_��V�>�ۇpc����Z�n�ӡ�A��d;��b�Pc֮�NT�_������N�����R�)����J�i},>�&�O���_,�ʜ     